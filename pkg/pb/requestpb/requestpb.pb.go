//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: requestpb/requestpb.proto

package requestpb

import (
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ReqNo    uint64 `protobuf:"varint,2,opt,name=req_no,json=reqNo,proto3" json:"req_no,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestpb_requestpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_requestpb_requestpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_requestpb_requestpb_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Request) GetReqNo() uint64 {
	if x != nil {
		return x.ReqNo
	}
	return 0
}

func (x *Request) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HashedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Req    *Request `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	Digest []byte   `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *HashedRequest) Reset() {
	*x = HashedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestpb_requestpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashedRequest) ProtoMessage() {}

func (x *HashedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requestpb_requestpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashedRequest.ProtoReflect.Descriptor instead.
func (*HashedRequest) Descriptor() ([]byte, []int) {
	return file_requestpb_requestpb_proto_rawDescGZIP(), []int{1}
}

func (x *HashedRequest) GetReq() *Request {
	if x != nil {
		return x.Req
	}
	return nil
}

func (x *HashedRequest) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*HashedRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestpb_requestpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_requestpb_requestpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_requestpb_requestpb_proto_rawDescGZIP(), []int{2}
}

func (x *Batch) GetRequests() []*HashedRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

var File_requestpb_requestpb_proto protoreflect.FileDescriptor

var file_requestpb_requestpb_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x22, 0x51, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x72, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x72, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4d, 0x0a, 0x0d, 0x48, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x72, 0x65,
	0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x69, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_requestpb_requestpb_proto_rawDescOnce sync.Once
	file_requestpb_requestpb_proto_rawDescData = file_requestpb_requestpb_proto_rawDesc
)

func file_requestpb_requestpb_proto_rawDescGZIP() []byte {
	file_requestpb_requestpb_proto_rawDescOnce.Do(func() {
		file_requestpb_requestpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_requestpb_requestpb_proto_rawDescData)
	})
	return file_requestpb_requestpb_proto_rawDescData
}

var file_requestpb_requestpb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_requestpb_requestpb_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: requestpb.Request
	(*HashedRequest)(nil), // 1: requestpb.HashedRequest
	(*Batch)(nil),         // 2: requestpb.Batch
}
var file_requestpb_requestpb_proto_depIdxs = []int32{
	0, // 0: requestpb.HashedRequest.req:type_name -> requestpb.Request
	1, // 1: requestpb.Batch.requests:type_name -> requestpb.HashedRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_requestpb_requestpb_proto_init() }
func file_requestpb_requestpb_proto_init() {
	if File_requestpb_requestpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_requestpb_requestpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_requestpb_requestpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashedRequest); i {
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
		file_requestpb_requestpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
			RawDescriptor: file_requestpb_requestpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_requestpb_requestpb_proto_goTypes,
		DependencyIndexes: file_requestpb_requestpb_proto_depIdxs,
		MessageInfos:      file_requestpb_requestpb_proto_msgTypes,
	}.Build()
	File_requestpb_requestpb_proto = out.File
	file_requestpb_requestpb_proto_rawDesc = nil
	file_requestpb_requestpb_proto_goTypes = nil
	file_requestpb_requestpb_proto_depIdxs = nil
}
