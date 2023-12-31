// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: calcservice.proto

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoMathClient is the client API for DoMath service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoMathClient interface {
	// ? Making a unary call for getting the calculation for two numbers
	CalcNumbers(ctx context.Context, in *OperationObj, opts ...grpc.CallOption) (*ResultObj, error)
	// ? Making a Stram client for the server calculation
	CalcListNumbers(ctx context.Context, opts ...grpc.CallOption) (DoMath_CalcListNumbersClient, error)
}

type doMathClient struct {
	cc grpc.ClientConnInterface
}

func NewDoMathClient(cc grpc.ClientConnInterface) DoMathClient {
	return &doMathClient{cc}
}

func (c *doMathClient) CalcNumbers(ctx context.Context, in *OperationObj, opts ...grpc.CallOption) (*ResultObj, error) {
	out := new(ResultObj)
	err := c.cc.Invoke(ctx, "/protofiles.DoMath/CalcNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doMathClient) CalcListNumbers(ctx context.Context, opts ...grpc.CallOption) (DoMath_CalcListNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &DoMath_ServiceDesc.Streams[0], "/protofiles.DoMath/CalcListNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &doMathCalcListNumbersClient{stream}
	return x, nil
}

type DoMath_CalcListNumbersClient interface {
	Send(*OperationObj) error
	CloseAndRecv() (*ResultObjList, error)
	grpc.ClientStream
}

type doMathCalcListNumbersClient struct {
	grpc.ClientStream
}

func (x *doMathCalcListNumbersClient) Send(m *OperationObj) error {
	return x.ClientStream.SendMsg(m)
}

func (x *doMathCalcListNumbersClient) CloseAndRecv() (*ResultObjList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResultObjList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoMathServer is the server API for DoMath service.
// All implementations must embed UnimplementedDoMathServer
// for forward compatibility
type DoMathServer interface {
	// ? Making a unary call for getting the calculation for two numbers
	CalcNumbers(context.Context, *OperationObj) (*ResultObj, error)
	// ? Making a Stram client for the server calculation
	CalcListNumbers(DoMath_CalcListNumbersServer) error
	mustEmbedUnimplementedDoMathServer()
}

// UnimplementedDoMathServer must be embedded to have forward compatible implementations.
type UnimplementedDoMathServer struct {
}

func (UnimplementedDoMathServer) CalcNumbers(context.Context, *OperationObj) (*ResultObj, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcNumbers not implemented")
}
func (UnimplementedDoMathServer) CalcListNumbers(DoMath_CalcListNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method CalcListNumbers not implemented")
}
func (UnimplementedDoMathServer) mustEmbedUnimplementedDoMathServer() {}

// UnsafeDoMathServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoMathServer will
// result in compilation errors.
type UnsafeDoMathServer interface {
	mustEmbedUnimplementedDoMathServer()
}

func RegisterDoMathServer(s grpc.ServiceRegistrar, srv DoMathServer) {
	s.RegisterService(&DoMath_ServiceDesc, srv)
}

func _DoMath_CalcNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationObj)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoMathServer).CalcNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.DoMath/CalcNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoMathServer).CalcNumbers(ctx, req.(*OperationObj))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoMath_CalcListNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DoMathServer).CalcListNumbers(&doMathCalcListNumbersServer{stream})
}

type DoMath_CalcListNumbersServer interface {
	SendAndClose(*ResultObjList) error
	Recv() (*OperationObj, error)
	grpc.ServerStream
}

type doMathCalcListNumbersServer struct {
	grpc.ServerStream
}

func (x *doMathCalcListNumbersServer) SendAndClose(m *ResultObjList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *doMathCalcListNumbersServer) Recv() (*OperationObj, error) {
	m := new(OperationObj)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoMath_ServiceDesc is the grpc.ServiceDesc for DoMath service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoMath_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.DoMath",
	HandlerType: (*DoMathServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalcNumbers",
			Handler:    _DoMath_CalcNumbers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalcListNumbers",
			Handler:       _DoMath_CalcListNumbers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calcservice.proto",
}
