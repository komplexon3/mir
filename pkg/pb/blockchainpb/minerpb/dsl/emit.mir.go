// Code generated by Mir codegen. DO NOT EDIT.

package minerpbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	events "github.com/filecoin-project/mir/pkg/pb/blockchainpb/minerpb/events"
	payloadpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/payloadpb"
	types "github.com/filecoin-project/mir/pkg/types"
)

// Module-specific dsl functions for emitting events.

func BlockRequest(m dsl.Module, destModule types.ModuleID, headId uint64, payload *payloadpb.Payload) {
	dsl.EmitMirEvent(m, events.BlockRequest(destModule, headId, payload))
}

func NewHead(m dsl.Module, destModule types.ModuleID, headId uint64) {
	dsl.EmitMirEvent(m, events.NewHead(destModule, headId))
}
