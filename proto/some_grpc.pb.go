// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/some.proto

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

const (
	SomeService_GetSome_FullMethodName = "/foo.SomeService/GetSome"
)

// SomeServiceClient is the client API for SomeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SomeServiceClient interface {
	GetSome(ctx context.Context, in *Some, opts ...grpc.CallOption) (*Some, error)
}

type someServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSomeServiceClient(cc grpc.ClientConnInterface) SomeServiceClient {
	return &someServiceClient{cc}
}

func (c *someServiceClient) GetSome(ctx context.Context, in *Some, opts ...grpc.CallOption) (*Some, error) {
	out := new(Some)
	err := c.cc.Invoke(ctx, SomeService_GetSome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SomeServiceServer is the server API for SomeService service.
// All implementations must embed UnimplementedSomeServiceServer
// for forward compatibility
type SomeServiceServer interface {
	GetSome(context.Context, *Some) (*Some, error)
	mustEmbedUnimplementedSomeServiceServer()
}

// UnimplementedSomeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSomeServiceServer struct {
}

func (UnimplementedSomeServiceServer) GetSome(context.Context, *Some) (*Some, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSome not implemented")
}
func (UnimplementedSomeServiceServer) mustEmbedUnimplementedSomeServiceServer() {}

// UnsafeSomeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SomeServiceServer will
// result in compilation errors.
type UnsafeSomeServiceServer interface {
	mustEmbedUnimplementedSomeServiceServer()
}

func RegisterSomeServiceServer(s grpc.ServiceRegistrar, srv SomeServiceServer) {
	s.RegisterService(&SomeService_ServiceDesc, srv)
}

func _SomeService_GetSome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Some)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomeServiceServer).GetSome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SomeService_GetSome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomeServiceServer).GetSome(ctx, req.(*Some))
	}
	return interceptor(ctx, in, info, handler)
}

// SomeService_ServiceDesc is the grpc.ServiceDesc for SomeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "foo.SomeService",
	HandlerType: (*SomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSome",
			Handler:    _SomeService_GetSome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/some.proto",
}
