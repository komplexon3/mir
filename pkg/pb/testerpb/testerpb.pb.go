// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: testerpb/testerpb.proto

package testerpb

import (
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

type Tester struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Tester) Reset() {
	*x = Tester{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testerpb_testerpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tester) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tester) ProtoMessage() {}

func (x *Tester) ProtoReflect() protoreflect.Message {
	mi := &file_testerpb_testerpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tester.ProtoReflect.Descriptor instead.
func (*Tester) Descriptor() ([]byte, []int) {
	return file_testerpb_testerpb_proto_rawDescGZIP(), []int{0}
}

var File_testerpb_testerpb_proto protoreflect.FileDescriptor

var file_testerpb_testerpb_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x70, 0x62, 0x1a, 0x1c, 0x6d, 0x69, 0x72, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0e, 0x0a, 0x06, 0x54, 0x65, 0x73, 0x74, 0x65, 0x72, 0x3a, 0x04, 0x98, 0xa6, 0x1d,
	0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_testerpb_testerpb_proto_rawDescOnce sync.Once
	file_testerpb_testerpb_proto_rawDescData = file_testerpb_testerpb_proto_rawDesc
)

func file_testerpb_testerpb_proto_rawDescGZIP() []byte {
	file_testerpb_testerpb_proto_rawDescOnce.Do(func() {
		file_testerpb_testerpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_testerpb_testerpb_proto_rawDescData)
	})
	return file_testerpb_testerpb_proto_rawDescData
}

var file_testerpb_testerpb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_testerpb_testerpb_proto_goTypes = []interface{}{
	(*Tester)(nil), // 0: testerpb.Tester
}
var file_testerpb_testerpb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testerpb_testerpb_proto_init() }
func file_testerpb_testerpb_proto_init() {
	if File_testerpb_testerpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testerpb_testerpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tester); i {
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
			RawDescriptor: file_testerpb_testerpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testerpb_testerpb_proto_goTypes,
		DependencyIndexes: file_testerpb_testerpb_proto_depIdxs,
		MessageInfos:      file_testerpb_testerpb_proto_msgTypes,
	}.Build()
	File_testerpb_testerpb_proto = out.File
	file_testerpb_testerpb_proto_rawDesc = nil
	file_testerpb_testerpb_proto_goTypes = nil
	file_testerpb_testerpb_proto_depIdxs = nil
}
