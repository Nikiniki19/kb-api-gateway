// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: engine.proto

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
	EngineRequest_FetchUser_FullMethodName = "/engine.EngineRequest/FetchUser"
)

// EngineRequestClient is the client API for EngineRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EngineRequestClient interface {
	FetchUser(ctx context.Context, in *EngineClientID, opts ...grpc.CallOption) (*EngineResponse, error)
}

type engineRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewEngineRequestClient(cc grpc.ClientConnInterface) EngineRequestClient {
	return &engineRequestClient{cc}
}

func (c *engineRequestClient) FetchUser(ctx context.Context, in *EngineClientID, opts ...grpc.CallOption) (*EngineResponse, error) {
	out := new(EngineResponse)
	err := c.cc.Invoke(ctx, EngineRequest_FetchUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineRequestServer is the server API for EngineRequest service.
// All implementations must embed UnimplementedEngineRequestServer
// for forward compatibility
type EngineRequestServer interface {
	FetchUser(context.Context, *EngineClientID) (*EngineResponse, error)
	mustEmbedUnimplementedEngineRequestServer()
}

// UnimplementedEngineRequestServer must be embedded to have forward compatible implementations.
type UnimplementedEngineRequestServer struct {
}

func (UnimplementedEngineRequestServer) FetchUser(context.Context, *EngineClientID) (*EngineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUser not implemented")
}
func (UnimplementedEngineRequestServer) mustEmbedUnimplementedEngineRequestServer() {}

// UnsafeEngineRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EngineRequestServer will
// result in compilation errors.
type UnsafeEngineRequestServer interface {
	mustEmbedUnimplementedEngineRequestServer()
}

func RegisterEngineRequestServer(s grpc.ServiceRegistrar, srv EngineRequestServer) {
	s.RegisterService(&EngineRequest_ServiceDesc, srv)
}

func _EngineRequest_FetchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EngineClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineRequestServer).FetchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EngineRequest_FetchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineRequestServer).FetchUser(ctx, req.(*EngineClientID))
	}
	return interceptor(ctx, in, info, handler)
}

// EngineRequest_ServiceDesc is the grpc.ServiceDesc for EngineRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EngineRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.EngineRequest",
	HandlerType: (*EngineRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchUser",
			Handler:    _EngineRequest_FetchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}
