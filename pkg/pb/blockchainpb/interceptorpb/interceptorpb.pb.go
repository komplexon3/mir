// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: blockchainpb/interceptorpb/interceptorpb.proto

package interceptorpb

import (
	blockchainpb "github.com/filecoin-project/mir/pkg/pb/blockchainpb"
	statepb "github.com/filecoin-project/mir/pkg/pb/blockchainpb/statepb"
	_ "github.com/filecoin-project/mir/pkg/pb/mir"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Event_TreeUpdate
	//	*Event_StateUpdate
	Type isEvent_Type `protobuf_oneof:"type"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescGZIP(), []int{0}
}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Event) GetTreeUpdate() *TreeUpdate {
	if x, ok := x.GetType().(*Event_TreeUpdate); ok {
		return x.TreeUpdate
	}
	return nil
}

func (x *Event) GetStateUpdate() *StateUpdate {
	if x, ok := x.GetType().(*Event_StateUpdate); ok {
		return x.StateUpdate
	}
	return nil
}

type isEvent_Type interface {
	isEvent_Type()
}

type Event_TreeUpdate struct {
	TreeUpdate *TreeUpdate `protobuf:"bytes,1,opt,name=tree_update,json=treeUpdate,proto3,oneof"`
}

type Event_StateUpdate struct {
	StateUpdate *StateUpdate `protobuf:"bytes,2,opt,name=state_update,json=stateUpdate,proto3,oneof"`
}

func (*Event_TreeUpdate) isEvent_Type() {}

func (*Event_StateUpdate) isEvent_Type() {}

type TreeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*blockchainpb.Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	HeadId uint64                `protobuf:"varint,2,opt,name=head_id,json=headId,proto3" json:"head_id,omitempty"`
}

func (x *TreeUpdate) Reset() {
	*x = TreeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreeUpdate) ProtoMessage() {}

func (x *TreeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreeUpdate.ProtoReflect.Descriptor instead.
func (*TreeUpdate) Descriptor() ([]byte, []int) {
	return file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescGZIP(), []int{1}
}

func (x *TreeUpdate) GetBlocks() []*blockchainpb.Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *TreeUpdate) GetHeadId() uint64 {
	if x != nil {
		return x.HeadId
	}
	return 0
}

type StateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *statepb.State `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StateUpdate) Reset() {
	*x = StateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateUpdate) ProtoMessage() {}

func (x *StateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateUpdate.ProtoReflect.Descriptor instead.
func (*StateUpdate) Descriptor() ([]byte, []int) {
	return file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescGZIP(), []int{2}
}

func (x *StateUpdate) GetState() *statepb.State {
	if x != nil {
		return x.State
	}
	return nil
}

var File_blockchainpb_interceptorpb_interceptorpb_proto protoreflect.FileDescriptor

var file_blockchainpb_interceptorpb_interceptorpb_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x1a,
	0x1f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65,
	0x6e, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x0b,
	0x74, 0x72, 0x65, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70,
	0x62, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0a,
	0x74, 0x72, 0x65, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x04, 0x90, 0xa6, 0x1d,
	0x01, 0x42, 0x0c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22,
	0x58, 0x0a, 0x0a, 0x54, 0x72, 0x65, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x49, 0x64, 0x3a, 0x04, 0x98, 0xa6, 0x1d, 0x01, 0x22, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x04,
	0x98, 0xa6, 0x1d, 0x01, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescOnce sync.Once
	file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescData = file_blockchainpb_interceptorpb_interceptorpb_proto_rawDesc
)

func file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescGZIP() []byte {
	file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescOnce.Do(func() {
		file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescData)
	})
	return file_blockchainpb_interceptorpb_interceptorpb_proto_rawDescData
}

var file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchainpb_interceptorpb_interceptorpb_proto_goTypes = []interface{}{
	(*Event)(nil),              // 0: interceptorpb.Event
	(*TreeUpdate)(nil),         // 1: interceptorpb.TreeUpdate
	(*StateUpdate)(nil),        // 2: interceptorpb.StateUpdate
	(*blockchainpb.Block)(nil), // 3: blockchainpb.Block
	(*statepb.State)(nil),      // 4: statepb.State
}
var file_blockchainpb_interceptorpb_interceptorpb_proto_depIdxs = []int32{
	1, // 0: interceptorpb.Event.tree_update:type_name -> interceptorpb.TreeUpdate
	2, // 1: interceptorpb.Event.state_update:type_name -> interceptorpb.StateUpdate
	3, // 2: interceptorpb.TreeUpdate.blocks:type_name -> blockchainpb.Block
	4, // 3: interceptorpb.StateUpdate.state:type_name -> statepb.State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blockchainpb_interceptorpb_interceptorpb_proto_init() }
func file_blockchainpb_interceptorpb_interceptorpb_proto_init() {
	if File_blockchainpb_interceptorpb_interceptorpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreeUpdate); i {
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
		file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateUpdate); i {
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
	file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_TreeUpdate)(nil),
		(*Event_StateUpdate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchainpb_interceptorpb_interceptorpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchainpb_interceptorpb_interceptorpb_proto_goTypes,
		DependencyIndexes: file_blockchainpb_interceptorpb_interceptorpb_proto_depIdxs,
		MessageInfos:      file_blockchainpb_interceptorpb_interceptorpb_proto_msgTypes,
	}.Build()
	File_blockchainpb_interceptorpb_interceptorpb_proto = out.File
	file_blockchainpb_interceptorpb_interceptorpb_proto_rawDesc = nil
	file_blockchainpb_interceptorpb_interceptorpb_proto_goTypes = nil
	file_blockchainpb_interceptorpb_interceptorpb_proto_depIdxs = nil
}
