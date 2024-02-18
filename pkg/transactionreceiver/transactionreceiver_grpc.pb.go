//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: transactionreceiver/transactionreceiver.proto

package transactionreceiver

import (
	context "context"
	trantorpb "github.com/filecoin-project/mir/pkg/pb/trantorpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TransactionReceiver_Listen_FullMethodName = "/receiver.TransactionReceiver/Listen"
)

// TransactionReceiverClient is the client API for TransactionReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionReceiverClient interface {
	Listen(ctx context.Context, opts ...grpc.CallOption) (TransactionReceiver_ListenClient, error)
}

type transactionReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionReceiverClient(cc grpc.ClientConnInterface) TransactionReceiverClient {
	return &transactionReceiverClient{cc}
}

func (c *transactionReceiverClient) Listen(ctx context.Context, opts ...grpc.CallOption) (TransactionReceiver_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &TransactionReceiver_ServiceDesc.Streams[0], TransactionReceiver_Listen_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transactionReceiverListenClient{stream}
	return x, nil
}

type TransactionReceiver_ListenClient interface {
	Send(*trantorpb.Transaction) error
	CloseAndRecv() (*ByeBye, error)
	grpc.ClientStream
}

type transactionReceiverListenClient struct {
	grpc.ClientStream
}

func (x *transactionReceiverListenClient) Send(m *trantorpb.Transaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transactionReceiverListenClient) CloseAndRecv() (*ByeBye, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ByeBye)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactionReceiverServer is the server API for TransactionReceiver service.
// All implementations must embed UnimplementedTransactionReceiverServer
// for forward compatibility
type TransactionReceiverServer interface {
	Listen(TransactionReceiver_ListenServer) error
	mustEmbedUnimplementedTransactionReceiverServer()
}

// UnimplementedTransactionReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionReceiverServer struct {
}

func (UnimplementedTransactionReceiverServer) Listen(TransactionReceiver_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedTransactionReceiverServer) mustEmbedUnimplementedTransactionReceiverServer() {}

// UnsafeTransactionReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionReceiverServer will
// result in compilation errors.
type UnsafeTransactionReceiverServer interface {
	mustEmbedUnimplementedTransactionReceiverServer()
}

func RegisterTransactionReceiverServer(s grpc.ServiceRegistrar, srv TransactionReceiverServer) {
	s.RegisterService(&TransactionReceiver_ServiceDesc, srv)
}

func _TransactionReceiver_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransactionReceiverServer).Listen(&transactionReceiverListenServer{stream})
}

type TransactionReceiver_ListenServer interface {
	SendAndClose(*ByeBye) error
	Recv() (*trantorpb.Transaction, error)
	grpc.ServerStream
}

type transactionReceiverListenServer struct {
	grpc.ServerStream
}

func (x *transactionReceiverListenServer) SendAndClose(m *ByeBye) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transactionReceiverListenServer) Recv() (*trantorpb.Transaction, error) {
	m := new(trantorpb.Transaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransactionReceiver_ServiceDesc is the grpc.ServiceDesc for TransactionReceiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionReceiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.TransactionReceiver",
	HandlerType: (*TransactionReceiverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _TransactionReceiver_Listen_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transactionreceiver/transactionreceiver.proto",
}
