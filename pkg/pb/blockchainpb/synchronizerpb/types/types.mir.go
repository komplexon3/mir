// Code generated by Mir codegen. DO NOT EDIT.

package synchronizerpbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types1 "github.com/filecoin-project/mir/codegen/model/types"
	blockchainpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	synchronizerpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/synchronizerpb"
	types "github.com/filecoin-project/mir/pkg/pb/blockchainpb/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Event struct {
	Type Event_Type
}

type Event_Type interface {
	mirreflect.GeneratedType
	isEvent_Type()
	Pb() synchronizerpb.Event_Type
}

type Event_TypeWrapper[T any] interface {
	Event_Type
	Unwrap() *T
}

func Event_TypeFromPb(pb synchronizerpb.Event_Type) Event_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *synchronizerpb.Event_SyncRequest:
		return &Event_SyncRequest{SyncRequest: SyncRequestFromPb(pb.SyncRequest)}
	}
	return nil
}

type Event_SyncRequest struct {
	SyncRequest *SyncRequest
}

func (*Event_SyncRequest) isEvent_Type() {}

func (w *Event_SyncRequest) Unwrap() *SyncRequest {
	return w.SyncRequest
}

func (w *Event_SyncRequest) Pb() synchronizerpb.Event_Type {
	if w == nil {
		return nil
	}
	if w.SyncRequest == nil {
		return &synchronizerpb.Event_SyncRequest{}
	}
	return &synchronizerpb.Event_SyncRequest{SyncRequest: (w.SyncRequest).Pb()}
}

func (*Event_SyncRequest) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.Event_SyncRequest]()}
}

func EventFromPb(pb *synchronizerpb.Event) *Event {
	if pb == nil {
		return nil
	}
	return &Event{
		Type: Event_TypeFromPb(pb.Type),
	}
}

func (m *Event) Pb() *synchronizerpb.Event {
	if m == nil {
		return nil
	}
	pbMessage := &synchronizerpb.Event{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Event) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.Event]()}
}

type SyncRequest struct {
	OrphanBlock  *types.Block
	LeaveNodeIds []uint64
}

func SyncRequestFromPb(pb *synchronizerpb.SyncRequest) *SyncRequest {
	if pb == nil {
		return nil
	}
	return &SyncRequest{
		OrphanBlock:  types.BlockFromPb(pb.OrphanBlock),
		LeaveNodeIds: pb.LeaveNodeIds,
	}
}

func (m *SyncRequest) Pb() *synchronizerpb.SyncRequest {
	if m == nil {
		return nil
	}
	pbMessage := &synchronizerpb.SyncRequest{}
	{
		if m.OrphanBlock != nil {
			pbMessage.OrphanBlock = (m.OrphanBlock).Pb()
		}
		pbMessage.LeaveNodeIds = m.LeaveNodeIds
	}

	return pbMessage
}

func (*SyncRequest) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.SyncRequest]()}
}

type Message struct {
	Type Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() synchronizerpb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb synchronizerpb.Message_Type) Message_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *synchronizerpb.Message_ChainRequest:
		return &Message_ChainRequest{ChainRequest: ChainRequestFromPb(pb.ChainRequest)}
	case *synchronizerpb.Message_ChainResponse:
		return &Message_ChainResponse{ChainResponse: ChainResponseFromPb(pb.ChainResponse)}
	}
	return nil
}

type Message_ChainRequest struct {
	ChainRequest *ChainRequest
}

func (*Message_ChainRequest) isMessage_Type() {}

func (w *Message_ChainRequest) Unwrap() *ChainRequest {
	return w.ChainRequest
}

func (w *Message_ChainRequest) Pb() synchronizerpb.Message_Type {
	if w == nil {
		return nil
	}
	if w.ChainRequest == nil {
		return &synchronizerpb.Message_ChainRequest{}
	}
	return &synchronizerpb.Message_ChainRequest{ChainRequest: (w.ChainRequest).Pb()}
}

func (*Message_ChainRequest) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.Message_ChainRequest]()}
}

type Message_ChainResponse struct {
	ChainResponse *ChainResponse
}

func (*Message_ChainResponse) isMessage_Type() {}

func (w *Message_ChainResponse) Unwrap() *ChainResponse {
	return w.ChainResponse
}

func (w *Message_ChainResponse) Pb() synchronizerpb.Message_Type {
	if w == nil {
		return nil
	}
	if w.ChainResponse == nil {
		return &synchronizerpb.Message_ChainResponse{}
	}
	return &synchronizerpb.Message_ChainResponse{ChainResponse: (w.ChainResponse).Pb()}
}

func (*Message_ChainResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.Message_ChainResponse]()}
}

func MessageFromPb(pb *synchronizerpb.Message) *Message {
	if pb == nil {
		return nil
	}
	return &Message{
		Type: Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *synchronizerpb.Message {
	if m == nil {
		return nil
	}
	pbMessage := &synchronizerpb.Message{}
	{
		if m.Type != nil {
			pbMessage.Type = (m.Type).Pb()
		}
	}

	return pbMessage
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.Message]()}
}

type ChainRequest struct {
	RequestId    string
	BlockId      uint64
	LeaveNodeIds []uint64
}

func ChainRequestFromPb(pb *synchronizerpb.ChainRequest) *ChainRequest {
	if pb == nil {
		return nil
	}
	return &ChainRequest{
		RequestId:    pb.RequestId,
		BlockId:      pb.BlockId,
		LeaveNodeIds: pb.LeaveNodeIds,
	}
}

func (m *ChainRequest) Pb() *synchronizerpb.ChainRequest {
	if m == nil {
		return nil
	}
	pbMessage := &synchronizerpb.ChainRequest{}
	{
		pbMessage.RequestId = m.RequestId
		pbMessage.BlockId = m.BlockId
		pbMessage.LeaveNodeIds = m.LeaveNodeIds
	}

	return pbMessage
}

func (*ChainRequest) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.ChainRequest]()}
}

type ChainResponse struct {
	RequestId string
	Found     bool
	Chain     []*types.Block
}

func ChainResponseFromPb(pb *synchronizerpb.ChainResponse) *ChainResponse {
	if pb == nil {
		return nil
	}
	return &ChainResponse{
		RequestId: pb.RequestId,
		Found:     pb.Found,
		Chain: types1.ConvertSlice(pb.Chain, func(t *blockchainpb.Block) *types.Block {
			return types.BlockFromPb(t)
		}),
	}
}

func (m *ChainResponse) Pb() *synchronizerpb.ChainResponse {
	if m == nil {
		return nil
	}
	pbMessage := &synchronizerpb.ChainResponse{}
	{
		pbMessage.RequestId = m.RequestId
		pbMessage.Found = m.Found
		pbMessage.Chain = types1.ConvertSlice(m.Chain, func(t *types.Block) *blockchainpb.Block {
			return (t).Pb()
		})
	}

	return pbMessage
}

func (*ChainResponse) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*synchronizerpb.ChainResponse]()}
}