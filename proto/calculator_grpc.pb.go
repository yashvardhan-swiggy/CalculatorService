// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/calculator.proto

package proto

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	GetSum(ctx context.Context, in *GetSumRequest, opts ...grpc.CallOption) (*GetSumResponse, error)
	GetPrimeNumbers(ctx context.Context, in *GetPrimeNumbersRequest, opts ...grpc.CallOption) (CalculatorService_GetPrimeNumbersClient, error)
	GetAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetAverageClient, error)
	GetMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetMaximumClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) GetSum(ctx context.Context, in *GetSumRequest, opts ...grpc.CallOption) (*GetSumResponse, error) {
	out := new(GetSumResponse)
	err := c.cc.Invoke(ctx, "/CalculatorService/GetSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) GetPrimeNumbers(ctx context.Context, in *GetPrimeNumbersRequest, opts ...grpc.CallOption) (CalculatorService_GetPrimeNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/CalculatorService/GetPrimeNumbers", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetPrimeNumbersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_GetPrimeNumbersClient interface {
	Recv() (*GetPrimeNumbersResponse, error)
	grpc.ClientStream
}

type calculatorServiceGetPrimeNumbersClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetPrimeNumbersClient) Recv() (*GetPrimeNumbersResponse, error) {
	m := new(GetPrimeNumbersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) GetAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/CalculatorService/GetAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetAverageClient{stream}
	return x, nil
}

type CalculatorService_GetAverageClient interface {
	Send(*GetAverageRequest) error
	CloseAndRecv() (*GetAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceGetAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetAverageClient) Send(m *GetAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceGetAverageClient) CloseAndRecv() (*GetAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GetAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) GetMaximum(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/CalculatorService/GetMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetMaximumClient{stream}
	return x, nil
}

type CalculatorService_GetMaximumClient interface {
	Send(*GetMaximumNumberRequest) error
	Recv() (*GetMaximumNumberResponse, error)
	grpc.ClientStream
}

type calculatorServiceGetMaximumClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetMaximumClient) Send(m *GetMaximumNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceGetMaximumClient) Recv() (*GetMaximumNumberResponse, error) {
	m := new(GetMaximumNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	GetSum(context.Context, *GetSumRequest) (*GetSumResponse, error)
	GetPrimeNumbers(*GetPrimeNumbersRequest, CalculatorService_GetPrimeNumbersServer) error
	GetAverage(CalculatorService_GetAverageServer) error
	GetMaximum(CalculatorService_GetMaximumServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) GetSum(context.Context, *GetSumRequest) (*GetSumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSum not implemented")
}
func (UnimplementedCalculatorServiceServer) GetPrimeNumbers(*GetPrimeNumbersRequest, CalculatorService_GetPrimeNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrimeNumbers not implemented")
}
func (UnimplementedCalculatorServiceServer) GetAverage(CalculatorService_GetAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) GetMaximum(CalculatorService_GetMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMaximum not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalculatorService/GetSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).GetSum(ctx, req.(*GetSumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_GetPrimeNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetPrimeNumbersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).GetPrimeNumbers(m, &calculatorServiceGetPrimeNumbersServer{stream})
}

type CalculatorService_GetPrimeNumbersServer interface {
	Send(*GetPrimeNumbersResponse) error
	grpc.ServerStream
}

type calculatorServiceGetPrimeNumbersServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetPrimeNumbersServer) Send(m *GetPrimeNumbersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_GetAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).GetAverage(&calculatorServiceGetAverageServer{stream})
}

type CalculatorService_GetAverageServer interface {
	SendAndClose(*GetAverageResponse) error
	Recv() (*GetAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceGetAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetAverageServer) SendAndClose(m *GetAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceGetAverageServer) Recv() (*GetAverageRequest, error) {
	m := new(GetAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_GetMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).GetMaximum(&calculatorServiceGetMaximumServer{stream})
}

type CalculatorService_GetMaximumServer interface {
	Send(*GetMaximumNumberResponse) error
	Recv() (*GetMaximumNumberRequest, error)
	grpc.ServerStream
}

type calculatorServiceGetMaximumServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetMaximumServer) Send(m *GetMaximumNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceGetMaximumServer) Recv() (*GetMaximumNumberRequest, error) {
	m := new(GetMaximumNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSum",
			Handler:    _CalculatorService_GetSum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrimeNumbers",
			Handler:       _CalculatorService_GetPrimeNumbers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAverage",
			Handler:       _CalculatorService_GetAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMaximum",
			Handler:       _CalculatorService_GetMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/calculator.proto",
}