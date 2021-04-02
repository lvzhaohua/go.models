// Code generated by protoc-gen-go-rpcx. DO NOT EDIT.

package tripurx_shopping_interface

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
// All implementations must embed UnimplementedDefaultRPCXServer
// for forward compatibility
type DefaultRPCXServer interface {
	//  GetSpu 获取Spu
	// 临时路径 tripurx_shopping.Default/GetSpu
	GetSpu(context.Context, *GetSpuRequest, *GetSpuReply) error
	mustEmbedUnimplementedDefaultRPCXServer()
}

// UnimplementedDefaultRPCXServer must be embedded to have forward compatible implementations.
type UnimplementedDefaultRPCXServer struct {
}

func (UnimplementedDefaultRPCXServer) GetSpu(context.Context, *GetSpuRequest, *GetSpuReply) error {
	return status.Errorf(codes.Unimplemented, "method GetSpu not implemented")
}
func (UnimplementedDefaultRPCXServer) mustEmbedUnimplementedDefaultRPCXServer() {}
