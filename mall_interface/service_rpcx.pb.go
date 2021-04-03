// Code generated by protoc-gen-go-rpcx. DO NOT EDIT.

package mall_interface

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

// DefaultRPCXServer is the server API for DefaultRPCX service.
// All implementations should embed UnimplementedDefaultRPCXServer
// for forward compatibility
type DefaultRPCXServer interface {
	Login(context.Context, *LoginRequest, *LoginReply) error
}

// UnimplementedDefaultRPCXServer should be embedded to have forward compatible implementations.
type UnimplementedDefaultRPCXServer struct {
}

func (UnimplementedDefaultRPCXServer) Login(context.Context, *LoginRequest, *LoginReply) error {
	return status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterDefaultRPCXServer(s grpc.ServiceRegistrar, srv DefaultRPCXServer) {
	s.RegisterService(&DefaultRPCX_ServiceDesc, srv)
}

// DefaultRPCX_ServiceDesc is the grpc.ServiceDesc for DefaultRPCX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DefaultRPCX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mall.Default",
	HandlerType: (*DefaultRPCXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall/service.proto",
}
