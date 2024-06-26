// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: proto/com.proto

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

// GeneralClient is the client API for General service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneralClient interface {
	SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Ip, error)
}

type generalClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneralClient(cc grpc.ClientConnInterface) GeneralClient {
	return &generalClient{cc}
}

func (c *generalClient) SendCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Ip, error) {
	out := new(Ip)
	err := c.cc.Invoke(ctx, "/code.General/sendCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneralServer is the server API for General service.
// All implementations must embed UnimplementedGeneralServer
// for forward compatibility
type GeneralServer interface {
	SendCommand(context.Context, *Command) (*Ip, error)
	mustEmbedUnimplementedGeneralServer()
}

// UnimplementedGeneralServer must be embedded to have forward compatible implementations.
type UnimplementedGeneralServer struct {
}

func (UnimplementedGeneralServer) SendCommand(context.Context, *Command) (*Ip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommand not implemented")
}
func (UnimplementedGeneralServer) mustEmbedUnimplementedGeneralServer() {}

// UnsafeGeneralServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneralServer will
// result in compilation errors.
type UnsafeGeneralServer interface {
	mustEmbedUnimplementedGeneralServer()
}

func RegisterGeneralServer(s grpc.ServiceRegistrar, srv GeneralServer) {
	s.RegisterService(&General_ServiceDesc, srv)
}

func _General_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/code.General/sendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralServer).SendCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

// General_ServiceDesc is the grpc.ServiceDesc for General service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var General_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "code.General",
	HandlerType: (*GeneralServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendCommand",
			Handler:    _General_SendCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/com.proto",
}
