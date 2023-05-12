package messagepbtypes

import (
	mirreflect "github.com/filecoin-project/mir/codegen/mirreflect"
	types3 "github.com/filecoin-project/mir/pkg/pb/availabilitypb/mscpb/types"
	types2 "github.com/filecoin-project/mir/pkg/pb/bcbpb/types"
	types4 "github.com/filecoin-project/mir/pkg/pb/checkpointpb/types"
	types1 "github.com/filecoin-project/mir/pkg/pb/isspb/types"
	messagepb "github.com/filecoin-project/mir/pkg/pb/messagepb"
	types5 "github.com/filecoin-project/mir/pkg/pb/ordererpb/types"
	pingpongpb "github.com/filecoin-project/mir/pkg/pb/pingpongpb"
	types "github.com/filecoin-project/mir/pkg/types"
	reflectutil "github.com/filecoin-project/mir/pkg/util/reflectutil"
)

type Message struct {
	DestModule types.ModuleID
	Type       Message_Type
}

type Message_Type interface {
	mirreflect.GeneratedType
	isMessage_Type()
	Pb() messagepb.Message_Type
}

type Message_TypeWrapper[T any] interface {
	Message_Type
	Unwrap() *T
}

func Message_TypeFromPb(pb messagepb.Message_Type) Message_Type {
	if pb == nil {
		return nil
	}
	switch pb := pb.(type) {
	case *messagepb.Message_Iss:
		return &Message_Iss{Iss: types1.ISSMessageFromPb(pb.Iss)}
	case *messagepb.Message_Bcb:
		return &Message_Bcb{Bcb: types2.MessageFromPb(pb.Bcb)}
	case *messagepb.Message_MultisigCollector:
		return &Message_MultisigCollector{MultisigCollector: types3.MessageFromPb(pb.MultisigCollector)}
	case *messagepb.Message_Pingpong:
		return &Message_Pingpong{Pingpong: pb.Pingpong}
	case *messagepb.Message_Checkpoint:
		return &Message_Checkpoint{Checkpoint: types4.MessageFromPb(pb.Checkpoint)}
	case *messagepb.Message_Orderer:
		return &Message_Orderer{Orderer: types5.MessageFromPb(pb.Orderer)}
	}
	return nil
}

type Message_Iss struct {
	Iss *types1.ISSMessage
}

func (*Message_Iss) isMessage_Type() {}

func (w *Message_Iss) Unwrap() *types1.ISSMessage {
	return w.Iss
}

func (w *Message_Iss) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_Iss{Iss: (w.Iss).Pb()}
}

func (*Message_Iss) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_Iss]()}
}

type Message_Bcb struct {
	Bcb *types2.Message
}

func (*Message_Bcb) isMessage_Type() {}

func (w *Message_Bcb) Unwrap() *types2.Message {
	return w.Bcb
}

func (w *Message_Bcb) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_Bcb{Bcb: (w.Bcb).Pb()}
}

func (*Message_Bcb) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_Bcb]()}
}

type Message_MultisigCollector struct {
	MultisigCollector *types3.Message
}

func (*Message_MultisigCollector) isMessage_Type() {}

func (w *Message_MultisigCollector) Unwrap() *types3.Message {
	return w.MultisigCollector
}

func (w *Message_MultisigCollector) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_MultisigCollector{MultisigCollector: (w.MultisigCollector).Pb()}
}

func (*Message_MultisigCollector) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_MultisigCollector]()}
}

type Message_Pingpong struct {
	Pingpong *pingpongpb.Message
}

func (*Message_Pingpong) isMessage_Type() {}

func (w *Message_Pingpong) Unwrap() *pingpongpb.Message {
	return w.Pingpong
}

func (w *Message_Pingpong) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_Pingpong{Pingpong: w.Pingpong}
}

func (*Message_Pingpong) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_Pingpong]()}
}

type Message_Checkpoint struct {
	Checkpoint *types4.Message
}

func (*Message_Checkpoint) isMessage_Type() {}

func (w *Message_Checkpoint) Unwrap() *types4.Message {
	return w.Checkpoint
}

func (w *Message_Checkpoint) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_Checkpoint{Checkpoint: (w.Checkpoint).Pb()}
}

func (*Message_Checkpoint) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_Checkpoint]()}
}

type Message_Orderer struct {
	Orderer *types5.Message
}

func (*Message_Orderer) isMessage_Type() {}

func (w *Message_Orderer) Unwrap() *types5.Message {
	return w.Orderer
}

func (w *Message_Orderer) Pb() messagepb.Message_Type {
	if w == nil {
		return nil
	}
	return &messagepb.Message_Orderer{Orderer: (w.Orderer).Pb()}
}

func (*Message_Orderer) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message_Orderer]()}
}

func MessageFromPb(pb *messagepb.Message) *Message {
	if pb == nil {
		return nil
	}
	return &Message{
		DestModule: (types.ModuleID)(pb.DestModule),
		Type:       Message_TypeFromPb(pb.Type),
	}
}

func (m *Message) Pb() *messagepb.Message {
	if m == nil {
		return nil
	}
	return &messagepb.Message{
		DestModule: (string)(m.DestModule),
		Type:       (m.Type).Pb(),
	}
}

func (*Message) MirReflect() mirreflect.Type {
	return mirreflect.TypeImpl{PbType_: reflectutil.TypeOf[*messagepb.Message]()}
}
