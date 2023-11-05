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
