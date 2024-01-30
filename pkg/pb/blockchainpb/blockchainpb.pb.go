// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: blockchainpb/blockchainpb.proto

package blockchainpb

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

type Blocktree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Leaves []uint64 `protobuf:"varint,2,rep,packed,name=leaves,proto3" json:"leaves,omitempty"`
}

func (x *Blocktree) Reset() {
	*x = Blocktree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_blockchainpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blocktree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blocktree) ProtoMessage() {}

func (x *Blocktree) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_blockchainpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blocktree.ProtoReflect.Descriptor instead.
func (*Blocktree) Descriptor() ([]byte, []int) {
	return file_blockchainpb_blockchainpb_proto_rawDescGZIP(), []int{0}
}

func (x *Blocktree) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *Blocktree) GetLeaves() []uint64 {
	if x != nil {
		return x.Leaves
	}
	return nil
}

type Blockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"` // ordered, no forks -> the 'current' chain
}

func (x *Blockchain) Reset() {
	*x = Blockchain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_blockchainpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blockchain) ProtoMessage() {}

func (x *Blockchain) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_blockchainpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blockchain.ProtoReflect.Descriptor instead.
func (*Blockchain) Descriptor() ([]byte, []int) {
	return file_blockchainpb_blockchainpb_proto_rawDescGZIP(), []int{1}
}

func (x *Blockchain) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId         uint64                 `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	PreviousBlockId uint64                 `protobuf:"varint,2,opt,name=previous_block_id,json=previousBlockId,proto3" json:"previous_block_id,omitempty"`
	Payload         *payloadpb.Payload     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchainpb_blockchainpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_blockchainpb_blockchainpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_blockchainpb_blockchainpb_proto_rawDescGZIP(), []int{2}
}

func (x *Block) GetBlockId() uint64 {
	if x != nil {
		return x.BlockId
	}
	return 0
}

func (x *Block) GetPreviousBlockId() uint64 {
	if x != nil {
		return x.PreviousBlockId
	}
	return 0
}

func (x *Block) GetPayload() *payloadpb.Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Block) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_blockchainpb_blockchainpb_proto protoreflect.FileDescriptor

var file_blockchainpb_blockchainpb_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x1a,
	0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x74,
	0x72, 0x65, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04,
	0x52, 0x06, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22, 0x3f,
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x06,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x22,
	0xbc, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a, 0x04, 0x80, 0xa6, 0x1d, 0x01, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69,
	0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchainpb_blockchainpb_proto_rawDescOnce sync.Once
	file_blockchainpb_blockchainpb_proto_rawDescData = file_blockchainpb_blockchainpb_proto_rawDesc
)

func file_blockchainpb_blockchainpb_proto_rawDescGZIP() []byte {
	file_blockchainpb_blockchainpb_proto_rawDescOnce.Do(func() {
		file_blockchainpb_blockchainpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchainpb_blockchainpb_proto_rawDescData)
	})
	return file_blockchainpb_blockchainpb_proto_rawDescData
}

var file_blockchainpb_blockchainpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_blockchainpb_blockchainpb_proto_goTypes = []interface{}{
	(*Blocktree)(nil),             // 0: blockchainpb.Blocktree
	(*Blockchain)(nil),            // 1: blockchainpb.Blockchain
	(*Block)(nil),                 // 2: blockchainpb.Block
	(*payloadpb.Payload)(nil),     // 3: payloadpb.Payload
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_blockchainpb_blockchainpb_proto_depIdxs = []int32{
	2, // 0: blockchainpb.Blocktree.blocks:type_name -> blockchainpb.Block
	2, // 1: blockchainpb.Blockchain.blocks:type_name -> blockchainpb.Block
	3, // 2: blockchainpb.Block.payload:type_name -> payloadpb.Payload
	4, // 3: blockchainpb.Block.timestamp:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_blockchainpb_blockchainpb_proto_init() }
func file_blockchainpb_blockchainpb_proto_init() {
	if File_blockchainpb_blockchainpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchainpb_blockchainpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blocktree); i {
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
		file_blockchainpb_blockchainpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blockchain); i {
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
		file_blockchainpb_blockchainpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
			RawDescriptor: file_blockchainpb_blockchainpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchainpb_blockchainpb_proto_goTypes,
		DependencyIndexes: file_blockchainpb_blockchainpb_proto_depIdxs,
		MessageInfos:      file_blockchainpb_blockchainpb_proto_msgTypes,
	}.Build()
	File_blockchainpb_blockchainpb_proto = out.File
	file_blockchainpb_blockchainpb_proto_rawDesc = nil
	file_blockchainpb_blockchainpb_proto_goTypes = nil
	file_blockchainpb_blockchainpb_proto_depIdxs = nil
}
