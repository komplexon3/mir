// Code generated by Mir codegen. DO NOT EDIT.

package interceptorpbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	blockchainpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	events "github.com/filecoin-project/mir/pkg/pb/blockchainpb/interceptorpb/events"
	types "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for emitting events.

func TreeUpdate(m dsl.Module, destModule types.ModuleID, tree *blockchainpb.Blocktree, headId uint64) {
	dsl.EmitMirEvent(m, events.TreeUpdate(destModule, tree, headId))
}

func NewOrphan(m dsl.Module, destModule types.ModuleID, orphan *blockchainpb.Block) {
	dsl.EmitMirEvent(m, events.NewOrphan(destModule, orphan))
}

func AppUpdate(m dsl.Module, destModule types.ModuleID, state int64) {
	dsl.EmitMirEvent(m, events.AppUpdate(destModule, state))
}