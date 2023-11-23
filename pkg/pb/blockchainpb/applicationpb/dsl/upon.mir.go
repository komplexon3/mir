// Code generated by Mir codegen. DO NOT EDIT.

package applicationpbdsl

import (
	dsl "github.com/filecoin-project/mir/pkg/dsl"
	blockchainpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	types "github.com/filecoin-project/mir/pkg/pb/blockchainpb/applicationpb/types"
	payloadpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/payloadpb"
	types1 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
)

// Module-specific dsl functions for processing events.

func UponEvent[W types.Event_TypeWrapper[Ev], Ev any](m dsl.Module, handler func(ev *Ev) error) {
	dsl.UponMirEvent[*types1.Event_Application](m, func(ev *types.Event) error {
		w, ok := ev.Type.(W)
		if !ok {
			return nil
		}

		return handler(w.Unwrap())
	})
}

func UponNewHead(m dsl.Module, handler func(headId uint64) error) {
	UponEvent[*types.Event_NewHead](m, func(ev *types.NewHead) error {
		return handler(ev.HeadId)
	})
}

func UponRegisterBlock(m dsl.Module, handler func(blockId *blockchainpb.Block) error) {
	UponEvent[*types.Event_RegisterBlock](m, func(ev *types.RegisterBlock) error {
		return handler(ev.BlockId)
	})
}

func UponPayloadRequest(m dsl.Module, handler func(headId uint64) error) {
	UponEvent[*types.Event_PayloadRequest](m, func(ev *types.PayloadRequest) error {
		return handler(ev.HeadId)
	})
}

func UponPayloadResponse(m dsl.Module, handler func(headId uint64, payload *payloadpb.Payload) error) {
	UponEvent[*types.Event_PayloadResponse](m, func(ev *types.PayloadResponse) error {
		return handler(ev.HeadId, ev.Payload)
	})
}
