// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mall_interface

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DefaultClient is the client API for Default service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DefaultClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type defaultClient struct {
	cc grpc.ClientConnInterface
}

func NewDefaultClient(cc grpc.ClientConnInterface) DefaultClient {
	return &defaultClient{cc}
}

func (c *defaultClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/mall.Default/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultServer is the server API for Default service.
// All implementations must embed UnimplementedDefaultServer
// for forward compatibility
type DefaultServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedDefaultServer()
}

// UnimplementedDefaultServer must be embedded to have forward compatible implementations.
type UnimplementedDefaultServer struct {
}

func (UnimplementedDefaultServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDefaultServer) mustEmbedUnimplementedDefaultServer() {}

// UnsafeDefaultServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DefaultServer will
// result in compilation errors.
type UnsafeDefaultServer interface {
	mustEmbedUnimplementedDefaultServer()
}

func RegisterDefaultServer(s grpc.ServiceRegistrar, srv DefaultServer) {
	s.RegisterService(&_Default_serviceDesc, srv)
}

func _Default_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.Default/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Default_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mall.Default",
	HandlerType: (*DefaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Default_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mall/service.proto",
}