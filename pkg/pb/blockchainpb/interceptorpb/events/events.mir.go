// Code generated by Mir codegen. DO NOT EDIT.

package interceptorpbevents

import (
	types3 "github.com/filecoin-project/mir/pkg/pb/blockchainpb/interceptorpb/types"
	types4 "github.com/filecoin-project/mir/pkg/pb/blockchainpb/statepb/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/blockchainpb/types"
	types2 "github.com/filecoin-project/mir/pkg/pb/eventpb/types"
	types "github.com/filecoin-project/mir/pkg/types"
)

func TreeUpdate(destModule types.ModuleID, blocks []*types1.Block, headId uint64) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Bcinterceptor{
			Bcinterceptor: &types3.Event{
				Type: &types3.Event_TreeUpdate{
					TreeUpdate: &types3.TreeUpdate{
						Blocks: blocks,
						HeadId: headId,
					},
				},
			},
		},
	}
}

func StateUpdate(destModule types.ModuleID, state *types4.State) *types2.Event {
	return &types2.Event{
		DestModule: destModule,
		Type: &types2.Event_Bcinterceptor{
			Bcinterceptor: &types3.Event{
				Type: &types3.Event_StateUpdate{
					StateUpdate: &types3.StateUpdate{
						State: state,
					},
				},
			},
		},
	}
}
