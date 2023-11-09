// Code generated by Mir codegen. DO NOT EDIT.

package bcmpbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	blockchainpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	events "github.com/filecoin-project/mir/pkg/pb/blockchainpb/bcmpb/events"
	types "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for emitting events.

func NewBlock(m dsl.Module, destModule types.ModuleID, block *blockchainpb.Block) {
	dsl.EmitMirEvent(m, events.NewBlock(destModule, block))
}

func NewChain(m dsl.Module, destModule types.ModuleID, blocks []*blockchainpb.Block) {
	dsl.EmitMirEvent(m, events.NewChain(destModule, blocks))
}

func GetBlockRequest(m dsl.Module, destModule types.ModuleID, requestId uint64, sourceModule types.ModuleID, blockId uint64) {
	dsl.EmitMirEvent(m, events.GetBlockRequest(destModule, requestId, sourceModule, blockId))
}

func GetBlockResponse(m dsl.Module, destModule types.ModuleID, requestId uint64, found bool, block *blockchainpb.Block) {
	dsl.EmitMirEvent(m, events.GetBlockResponse(destModule, requestId, found, block))
}
