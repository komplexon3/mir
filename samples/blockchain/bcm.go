// blockchain manager

package main

import (
	"errors"
	"fmt"

	"github.com/filecoin-project/mir/pkg/dsl"
	"github.com/filecoin-project/mir/pkg/logging"
	"github.com/filecoin-project/mir/pkg/modules"
	"github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	applicationpbdsl "github.com/filecoin-project/mir/pkg/pb/blockchainpb/applicationpb/dsl"
	bcmpbdsl "github.com/filecoin-project/mir/pkg/pb/blockchainpb/bcmpb/dsl"
	interceptorpbdsl "github.com/filecoin-project/mir/pkg/pb/blockchainpb/interceptorpb/dsl"
	minerpbdsl "github.com/filecoin-project/mir/pkg/pb/blockchainpb/minerpb/dsl"
	"github.com/filecoin-project/mir/pkg/pb/blockchainpb/payloadpb"
	"github.com/filecoin-project/mir/pkg/pb/blockchainpb/statepb"
	synchronizerpbdsl "github.com/filecoin-project/mir/pkg/pb/blockchainpb/synchronizerpb/dsl"
	t "github.com/filecoin-project/mir/pkg/types"
	"github.com/filecoin-project/mir/samples/blockchain/utils"
	"golang.org/x/exp/maps"
)

var (
	ErrParentNotFound    = errors.New("parent not found")
	ErrNodeAlreadyInTree = errors.New("node already in tree")
)

type bcmBlock struct {
	block     *blockchainpb.BlockInternal
	parent    *bcmBlock
	depth     uint64
	scanCount uint64
}

type bcmModule struct {
	m           *dsl.Module
	blocks      []bcmBlock // IMPORTANT: only to be used for visualization, not for actual block lookup
	leaves      map[uint64]*bcmBlock
	head        *bcmBlock
	genesis     *bcmBlock
	checkpoints map[uint64]*bcmBlock
	blockCount  uint64
	logger      logging.Logger
}

func (bcm *bcmModule) handleNewHead(newHead *blockchainpb.BlockInternal) {
	bcm.logger.Log(logging.LevelInfo, "Head changed", "head", utils.FormatBlockId(newHead.Block.BlockId))
	applicationpbdsl.NewHead(*bcm.m, "application", newHead.Block.BlockId)
	minerpbdsl.NewHead(*bcm.m, "miner", newHead.Block.BlockId)
}

// traversalFunc should return true if traversal should continue, false otherwise
// if it shouldn't continue, the function will return the current block
// if it should continue but there are no more blocks, the function will return nil
// if an error is thrown, the traversal is aborted and nil is returned
func (bcm *bcmModule) blockTreeTraversal(traversalFunc func(currBlock *bcmBlock) (bool, error)) *bcmBlock {
	currentScanCount := uint64(0) // to mark nodes that have been scanned in this round
	queue := make([]*bcmBlock, 0)
	for _, v := range bcm.leaves {
		queue = append(queue, v)
		if v.scanCount > currentScanCount {
			currentScanCount = v.scanCount
		}
	}
	for len(queue) > 0 {
		// pop from queue
		curr := queue[0]
		queue = queue[1:]

		// mark as scanned
		curr.scanCount = currentScanCount + 1

		if match, err := traversalFunc(curr); err != nil {
			return nil
		} else if match {
			return curr
		}

		// add curr's parents to queue if it hasn't been scanned yet
		if curr.parent != nil && curr.parent.scanCount <= currentScanCount {
			queue = append(queue, curr.parent)
		}
	}
	return nil
}

// TODO: might need to add some "cleanup" to handle very old leaves in case this grows too much
func (bcm *bcmModule) addBlock(block *blockchainpb.Block) error {
	bcm.logger.Log(logging.LevelInfo, "Adding block...", "blockId", utils.FormatBlockId(block.BlockId))

	// check if block is already in the leaves, reject if so
	if _, ok := bcm.leaves[block.BlockId]; ok {
		bcm.logger.Log(logging.LevelDebug, "Block already in leaves - ignore", "blockId", utils.FormatBlockId(block.BlockId))
		return nil
	}

	parentId := block.PreviousBlockId
	// check leaves for parent with its id
	// return nil if it isn't there
	if parent, ok := bcm.leaves[parentId]; ok {
		bcm.logger.Log(logging.LevelDebug, "Found parend in leaves", "blockId", utils.FormatBlockId(block.BlockId), "parentId", utils.FormatBlockId(parentId))
		blockNode := bcmBlock{
			block: &blockchainpb.BlockInternal{
				Block: block,
				State: nil,
			},
			parent: parent,
			depth:  parent.depth + 1,
		}
		bcm.blocks = append(bcm.blocks, blockNode)
		// replace leave
		delete(bcm.leaves, parentId)
		bcm.leaves[block.BlockId] = &blockNode
		if blockNode.depth > bcm.head.depth { // update head only if new block is deeper, i.e. is the longest chain
			bcm.head = &blockNode
		}
		bcm.blockCount++
		return nil
	}

	if hit := bcm.blockTreeTraversal(func(currBlock *bcmBlock) (bool, error) {
		// check if curr matches the block to be added - if so, ignore
		if currBlock.block.Block.BlockId == block.BlockId {
			bcm.logger.Log(logging.LevelDebug, "Block already in tree", "blockId", utils.FormatBlockId(block.BlockId))
			return true, ErrNodeAlreadyInTree
		}

		// check if curr is parent
		if currBlock.block.Block.BlockId == parentId {
			bcm.logger.Log(logging.LevelDebug, "Found parend in tree", "blockId", utils.FormatBlockId(block.BlockId), "parentId", utils.FormatBlockId(parentId))
			blockNode := bcmBlock{
				block: &blockchainpb.BlockInternal{
					Block: block,
					State: nil,
				},
				parent:    currBlock,
				depth:     currBlock.depth + 1,
				scanCount: 0,
			}
			// add new leave
			bcm.leaves[block.BlockId] = &blockNode
			bcm.blocks = append(bcm.blocks, blockNode)
			if blockNode.depth > bcm.head.depth { // update head only if new block is deeper, i.e. is the longest chain
				bcm.head = &blockNode
			}
			bcm.blockCount++
			return true, nil
		}

		return false, nil
	}); hit == nil {
		bcm.logger.Log(logging.LevelInfo, "Couldn't find parent", "blockId", utils.FormatBlockId(block.BlockId))
		return ErrParentNotFound
	}

	return nil
}

func (bcm *bcmModule) findBlock(blockId uint64) *bcmBlock {
	return bcm.blockTreeTraversal(func(currBlock *bcmBlock) (bool, error) {
		if currBlock.block.Block.BlockId == blockId {
			return true, nil
		}
		return false, nil
	})
}

func (bcm *bcmModule) getHead() *blockchainpb.BlockInternal {
	return bcm.head.block
}

func (bcm *bcmModule) handleNewBlock(block *blockchainpb.Block) {
	currentHead := bcm.getHead()
	// insert block
	if err := bcm.addBlock(block); err != nil {
		switch err {
		case ErrNodeAlreadyInTree:
			// ignore
			bcm.logger.Log(logging.LevelDebug, "Block already in tree - ignore", "blockId", utils.FormatBlockId(block.BlockId))
		case ErrParentNotFound:
			// ask synchronizer for help
			leavesPlusGenesis := append(maps.Keys(bcm.leaves), bcm.genesis.block.Block.BlockId)
			bcm.logger.Log(logging.LevelDebug, "Sending sync request", "blockId", utils.FormatBlockId(block.BlockId), "leaves+genesis", utils.FormatBlockIdSlice(leavesPlusGenesis))
			synchronizerpbdsl.SyncRequest(*bcm.m, "synchronizer", block, leavesPlusGenesis)
			// announce orphan block to interceptor, just for the visualization
			interceptorpbdsl.NewOrphan(*bcm.m, "devnull", block)
		default:
			// some other error
			bcm.logger.Log(logging.LevelError, "Unexpected error adding block", "blockId", utils.FormatBlockId(block.BlockId), "error", err)
			panic(err)
		}
		return
	}

	// if head changed, trigger get tpm to prep new block
	if newHead := bcm.getHead(); newHead != currentHead {
		bcm.handleNewHead(newHead)
	}

	// send to chainviewer
	bcm.sendTreeUpdate()
}

func (bcm *bcmModule) handleNewChain(blocks []*blockchainpb.Block) {
	currentHead := bcm.getHead()
	blockIds := make([]uint64, 0, len(blocks))
	for _, v := range blocks {
		blockIds = append(blockIds, v.BlockId)
	}
	bcm.logger.Log(logging.LevelDebug, "Adding new chain", "chain", utils.FormatBlockIdSlice(blockIds))
	// insert block
	for _, block := range blocks {
		if err := bcm.addBlock(block); err != nil {
			switch err {
			case ErrNodeAlreadyInTree:
				// ignore
				bcm.logger.Log(logging.LevelDebug, "Block already in tree - ignore", "blockId", utils.FormatBlockId(block.BlockId))
			case ErrParentNotFound:
				// ask synchronizer for help
				leavesPlusGenesis := append(maps.Keys(bcm.leaves), bcm.genesis.block.Block.BlockId)
				bcm.logger.Log(logging.LevelDebug, "Sending sync request", "blockId", utils.FormatBlockId(block.BlockId), "leaves+genesis", utils.FormatBlockIdSlice(leavesPlusGenesis))
				synchronizerpbdsl.SyncRequest(*bcm.m, "synchronizer", block, leavesPlusGenesis)
				// announce orphan block to interceptor
				interceptorpbdsl.NewOrphan(*bcm.m, "devnull", block)
			default:
				// some other error
				bcm.logger.Log(logging.LevelError, "Unexpected error adding block", "blockId", utils.FormatBlockId(block.BlockId), "error", err)
				panic(err)
			}
		}
	}

	// if head changed, trigger get tpm to prep new block (only need to do this once)
	if newHead := bcm.getHead(); newHead != currentHead {
		bcm.handleNewHead(newHead)
	}

	// send to chainviewer
	bcm.sendTreeUpdate()
}

func (bcm *bcmModule) sendTreeUpdate() {
	blocks := func() []*blockchainpb.Block {
		blocks := make([]*blockchainpb.Block, 0, len(bcm.blocks))
		for _, v := range bcm.blocks {
			blocks = append(blocks, v.block.Block)
		}
		return blocks
	}()
	leaves := func() []uint64 {
		leaves := make([]uint64, 0, len(bcm.leaves))
		for _, v := range bcm.leaves {
			leaves = append(leaves, v.block.Block.BlockId)
		}
		return leaves
	}()

	blockTree := blockchainpb.Blocktree{Blocks: blocks, Leaves: leaves}

	interceptorpbdsl.TreeUpdate(*bcm.m, "devnull", &blockTree, bcm.head.block.Block.BlockId)
}

func formatBlockInts(blocks []*blockchainpb.BlockInternal) []string {
	blockIds := make([]string, 0, len(blocks))
	for _, v := range blocks {
		stateText := "no state"
		if v.State != nil {
			stateText = v.State.String()
		}
		blockIds = append(blockIds, fmt.Sprintf("%d - %s", v.Block.BlockId, stateText))
	}
	return blockIds
}

func (bcm *bcmModule) handleGetHeadToCheckpointChainRequest(requestID string, sourceModule t.ModuleID) error {
	// for debugging
	// TODO: remove
	checkpoints := make([]*blockchainpb.BlockInternal, 0, len(bcm.checkpoints))

	for _, value := range bcm.checkpoints {
		checkpoints = append(checkpoints, value.block)
	}

	bcm.logger.Log(logging.LevelInfo, "Received get head to checkpoint chain request", "requestId", requestID, "checkpoints", formatBlockInts(checkpoints))
	chain := make([]*blockchainpb.BlockInternal, 0)

	// start with head
	currentBlock := bcm.head
	for currentBlock != nil {
		if cp, ok := bcm.checkpoints[currentBlock.block.Block.BlockId]; ok {
			chain = append(chain, cp.block)
			break
		}
		chain = append(chain, currentBlock.block)
		currentBlock = currentBlock.parent
	}

	if currentBlock == nil {
		bcm.logger.Log(logging.LevelError, "Cannot link up with checkpoints - this should not happen", "chain", chain)
		panic("Cannot link up with checkpoints - this should not happen")
	}

	for i, j := 0, len(chain)-1; i < j; i, j = i+1, j-1 {
		chain[i], chain[j] = chain[j], chain[i]
	}

	bcm.logger.Log(logging.LevelDebug, "Sending chain for state comp", "requestId", requestID, "chain", formatBlockInts(chain), "checkpoint", formatBlockInts([]*blockchainpb.BlockInternal{currentBlock.block}))
	bcmpbdsl.GetHeadToCheckpointChainResponse(*bcm.m, sourceModule, requestID, chain)

	return nil

}

func (bcm *bcmModule) handleRegisterCheckpoint(block_id uint64, state *statepb.State) error {
	bcm.logger.Log(logging.LevelInfo, "Received register checkpoint", "blockId", utils.FormatBlockId(block_id))
	if _, ok := bcm.checkpoints[block_id]; ok {
		bcm.logger.Log(logging.LevelDebug, "Checkpoint already registered - ignore", "blockId", utils.FormatBlockId(block_id))
		return nil
	}

	// find block
	// should be a leave or a block really close
	if hit := bcm.findBlock(block_id); hit == nil {
		bcm.logger.Log(logging.LevelError, "Block not found - can't register checkpoint", "blockId", utils.FormatBlockId(block_id))
		panic("block not found when registering checkpoint") // should never happen
	} else {
		bcm.logger.Log(logging.LevelDebug, "Found block - registering checkpoint", "blockId", utils.FormatBlockId(block_id))
		hit.block.State = state
		bcm.checkpoints[block_id] = hit
	}

	return nil
}

func NewBCM(logger logging.Logger) modules.PassiveModule {
	m := dsl.NewModule("bcm")

	// making up a genisis block
	genesis := &blockchainpb.Block{
		BlockId:         0,
		PreviousBlockId: 0,
		Payload:         &payloadpb.Payload{AddMinus: 0},
		Timestamp:       0, // unix 0
	}

	hash := utils.HashBlock(genesis)
	genesis.BlockId = hash
	genesisBcm := bcmBlock{
		block: &blockchainpb.BlockInternal{
			Block: genesis,
			State: &statepb.State{
				Counter: 0,
			},
		},
		parent: nil,
		depth:  0,
	}

	// init bcm genisis block
	bcm := bcmModule{
		m:           &m,
		blocks:      []bcmBlock{genesisBcm},
		leaves:      make(map[uint64]*bcmBlock),
		head:        &genesisBcm,
		genesis:     &genesisBcm,
		checkpoints: make(map[uint64]*bcmBlock),
		blockCount:  1,
		logger:      logger,
	}

	// add genesis to leaves and checkpoints
	bcm.leaves[hash] = bcm.head
	bcm.checkpoints[hash] = bcm.head

	dsl.UponInit(m, func() error {

		// start with genesis block
		applicationpbdsl.NewHead(m, "application", hash)
		// start mining on genesis block
		minerpbdsl.NewHead(m, "miner", hash)

		return nil
	})

	bcmpbdsl.UponNewBlock(m, func(block *blockchainpb.Block) error {
		bcm.handleNewBlock(block)
		return nil
	})

	bcmpbdsl.UponNewChain(m, func(blocks []*blockchainpb.Block) error {
		bcm.logger.Log(logging.LevelInfo, "Received chain from synchronizer")
		bcm.handleNewChain(blocks)
		return nil
	})

	bcmpbdsl.UponGetBlockRequest(m, func(requestID string, sourceModule t.ModuleID, blockID uint64) error {
		bcm.logger.Log(logging.LevelInfo, "Received get block request", "requestId", requestID, "sourceModule", sourceModule)
		// check if block is in tree
		hit := bcm.findBlock(blockID)

		if hit == nil {
			bcm.logger.Log(logging.LevelDebug, "Block not found", "requestId", requestID)
			bcmpbdsl.GetBlockResponse(*bcm.m, sourceModule, requestID, false, nil)
			return nil
		}

		bcm.logger.Log(logging.LevelDebug, "Found block in tree", "requestId", requestID)
		bcmpbdsl.GetBlockResponse(*bcm.m, sourceModule, requestID, true, hit.block.GetBlock())

		return nil
	})

	bcmpbdsl.UponGetChainRequest(m, func(requestID string, sourceModule t.ModuleID, endBlockId uint64, sourceBlockIds []uint64) error {
		bcm.logger.Log(logging.LevelInfo, "Received get chain request", "requestId", requestID, "sourceModule", sourceModule, "endBlockId", utils.FormatBlockId(endBlockId), "sourceBlockIds", utils.FormatBlockIdSlice(sourceBlockIds))
		chain := make([]*blockchainpb.Block, 0)
		// for easier lookup...
		sourceBlockIdsMap := make(map[uint64]uint64)
		for _, v := range sourceBlockIds {
			sourceBlockIdsMap[v] = v
		}
		currentBlock := bcm.findBlock(endBlockId)

		if currentBlock == nil {
			bcm.logger.Log(logging.LevelDebug, "Block not found - can't build chain", "requestId", requestID)
			bcmpbdsl.GetBlockResponse(*bcm.m, sourceModule, requestID, false, nil)
			return nil
		}

		bcm.logger.Log(logging.LevelDebug, "Found block in tree - building chain", "requestId", requestID)

		// initialize chain
		chain = append(chain, currentBlock.block.Block)

		for currentBlock.parent != nil {
			// if parent is in sourceBlockIds, stop
			if _, ok := sourceBlockIdsMap[currentBlock.parent.block.Block.BlockId]; ok {
				// chain build
				bcm.logger.Log(logging.LevelDebug, "Found link with sourceBlockIds - chain build", "requestId", requestID)
				// reversing chain to send it in the right order
				for i, j := 0, len(chain)-1; i < j; i, j = i+1, j-1 {
					chain[i], chain[j] = chain[j], chain[i]
				}
				bcmpbdsl.GetChainResponse(*bcm.m, sourceModule, requestID, true, chain)
				return nil
			}

			chain = append(chain, currentBlock.parent.block.Block)
			currentBlock = currentBlock.parent
		}

		// if we get here, we didn't find a link with sourceBlockIds
		bcm.logger.Log(logging.LevelDebug, "No link with sourceBlockIds - can't build chain", "requestId", requestID)
		bcmpbdsl.GetChainResponse(*bcm.m, sourceModule, requestID, false, nil)

		return nil

	})

	bcmpbdsl.UponGetHeadToCheckpointChainRequest(m, bcm.handleGetHeadToCheckpointChainRequest)
	bcmpbdsl.UponRegisterCheckpoint(m, bcm.handleRegisterCheckpoint)

	return m
}