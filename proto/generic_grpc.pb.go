// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: generic.proto

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
	GenericRequest_FetchUser_FullMethodName = "/generic.GenericRequest/FetchUser"
)

// GenericRequestClient is the client API for GenericRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenericRequestClient interface {
	FetchUser(ctx context.Context, in *GenericClientID, opts ...grpc.CallOption) (*GenericResponse, error)
}

type genericRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericRequestClient(cc grpc.ClientConnInterface) GenericRequestClient {
	return &genericRequestClient{cc}
}

func (c *genericRequestClient) FetchUser(ctx context.Context, in *GenericClientID, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, GenericRequest_FetchUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericRequestServer is the server API for GenericRequest service.
// All implementations must embed UnimplementedGenericRequestServer
// for forward compatibility
type GenericRequestServer interface {
	FetchUser(context.Context, *GenericClientID) (*GenericResponse, error)
	mustEmbedUnimplementedGenericRequestServer()
}

// UnimplementedGenericRequestServer must be embedded to have forward compatible implementations.
type UnimplementedGenericRequestServer struct {
}

func (UnimplementedGenericRequestServer) FetchUser(context.Context, *GenericClientID) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUser not implemented")
}
func (UnimplementedGenericRequestServer) mustEmbedUnimplementedGenericRequestServer() {}

// UnsafeGenericRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenericRequestServer will
// result in compilation errors.
type UnsafeGenericRequestServer interface {
	mustEmbedUnimplementedGenericRequestServer()
}

func RegisterGenericRequestServer(s grpc.ServiceRegistrar, srv GenericRequestServer) {
	s.RegisterService(&GenericRequest_ServiceDesc, srv)
}

func _GenericRequest_FetchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericRequestServer).FetchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenericRequest_FetchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericRequestServer).FetchUser(ctx, req.(*GenericClientID))
	}
	return interceptor(ctx, in, info, handler)
}

// GenericRequest_ServiceDesc is the grpc.ServiceDesc for GenericRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenericRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generic.GenericRequest",
	HandlerType: (*GenericRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchUser",
			Handler:    _GenericRequest_FetchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generic.proto",
}
