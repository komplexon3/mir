// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: blockchainpb/statepb/statepb.proto

package statepb

import (
	payloadpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/payloadpb"
	_ "github.com/filecoin-project/mir/pkg/pb/mir"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// application specific state
	History            []*payloadpb.Payload `protobuf:"bytes,1,rep,name=history,proto3" json:"history,omitempty"`
	LastSentTimestamps []*LastSentTimestamp `protobuf:"bytes,2,rep,name=last_sent_timestamps,json=lastSentTimestamps,proto3" json:"last_sent_timestamps,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_statepb_statepb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_statepb_statepb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_blockchainpb_statepb_statepb_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetHistory() []*payloadpb.Payload {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *State) GetLastSentTimestamps() []*LastSentTimestamp {
	if x != nil {
		return x.LastSentTimestamps
	}
	return nil
}

type LastSentTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId    string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LastSentTimestamp) Reset() {
	*x = LastSentTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_statepb_statepb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastSentTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastSentTimestamp) ProtoMessage() {}

func (x *LastSentTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_statepb_statepb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastSentTimestamp.ProtoReflect.Descriptor instead.
func (*LastSentTimestamp) Descriptor() ([]byte, []int) {
	return file_blockchainpb_statepb_statepb_proto_rawDescGZIP(), []int{1}
}

func (x *LastSentTimestamp) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *LastSentTimestamp) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_blockchainpb_statepb_statepb_proto protoreflect.FileDescriptor

var file_blockchainpb_statepb_statepb_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x1a, 0x26, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67,
	0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x4c, 0x0a, 0x14,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x70, 0x62, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01,
	0x22, 0xa2, 0x01, 0x0a, 0x11, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4d, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0x82, 0xa6, 0x1d, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e,
	0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a,
	0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchainpb_statepb_statepb_proto_rawDescOnce sync.Once
	file_blockchainpb_statepb_statepb_proto_rawDescData = file_blockchainpb_statepb_statepb_proto_rawDesc
)

func file_blockchainpb_statepb_statepb_proto_rawDescGZIP() []byte {
	file_blockchainpb_statepb_statepb_proto_rawDescOnce.Do(func() {
		file_blockchainpb_statepb_statepb_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchainpb_statepb_statepb_proto_rawDescData)
	})
	return file_blockchainpb_statepb_statepb_proto_rawDescData
}

var file_blockchainpb_statepb_statepb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blockchainpb_statepb_statepb_proto_goTypes = []interface{}{
	(*State)(nil),                 // 0: statepb.State
	(*LastSentTimestamp)(nil),     // 1: statepb.LastSentTimestamp
	(*payloadpb.Payload)(nil),     // 2: payloadpb.Payload
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_blockchainpb_statepb_statepb_proto_depIdxs = []int32{
	2, // 0: statepb.State.history:type_name -> payloadpb.Payload
	1, // 1: statepb.State.last_sent_timestamps:type_name -> statepb.LastSentTimestamp
	3, // 2: statepb.LastSentTimestamp.timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blockchainpb_statepb_statepb_proto_init() }
func file_blockchainpb_statepb_statepb_proto_init() {
	if File_blockchainpb_statepb_statepb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchainpb_statepb_statepb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blockchainpb_statepb_statepb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastSentTimestamp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchainpb_statepb_statepb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchainpb_statepb_statepb_proto_goTypes,
		DependencyIndexes: file_blockchainpb_statepb_statepb_proto_depIdxs,
		MessageInfos:      file_blockchainpb_statepb_statepb_proto_msgTypes,
	}.Build()
	File_blockchainpb_statepb_statepb_proto = out.File
	file_blockchainpb_statepb_statepb_proto_rawDesc = nil
	file_blockchainpb_statepb_statepb_proto_goTypes = nil
	file_blockchainpb_statepb_statepb_proto_depIdxs = nil
}
