// Code generated by Mir codegen. DO NOT EDIT.

package payloadpbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	payloadpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/payloadpb"
	types "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Payload struct {
	Message   string
	Timestamp *timestamppb.Timestamp
	Sender    types.NodeID
}

func PayloadFromPb(pb *payloadpb.Payload) *Payload {
	if pb == nil {
		return nil
	}
	return &Payload{
		Message:   pb.Message,
		Timestamp: pb.Timestamp,
		Sender:    (types.NodeID)(pb.Sender),
	}
}

func (m *Payload) Pb() *payloadpb.Payload {
	if m == nil {
		return nil
	}
	pbMessage := &payloadpb.Payload{}
	{
		pbMessage.Message = m.Message
		if m.Timestamp != nil {
			pbMessage.Timestamp = m.Timestamp
		}
		pbMessage.Sender = (string)(m.Sender)
	}

	return pbMessage
}

func (*Payload) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*payloadpb.Payload]()}
}