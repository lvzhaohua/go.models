// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package novel_spider_interface

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
	// SearchNovel 搜索小说
	SearchNovel(ctx context.Context, in *SearchNovelRequest, opts ...grpc.CallOption) (*SearchNovelReply, error)
	// SearchMenu 搜索目录
	SearchMenu(ctx context.Context, in *SearchMenuRequest, opts ...grpc.CallOption) (*SearchMenuReply, error)
	// SearchChapter 搜索章节
	SearchChapter(ctx context.Context, in *SearchChapterRequest, opts ...grpc.CallOption) (*SearchChapterReply, error)
}

type defaultClient struct {
	cc grpc.ClientConnInterface
}

func NewDefaultClient(cc grpc.ClientConnInterface) DefaultClient {
	return &defaultClient{cc}
}

func (c *defaultClient) SearchNovel(ctx context.Context, in *SearchNovelRequest, opts ...grpc.CallOption) (*SearchNovelReply, error) {
	out := new(SearchNovelReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchNovel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultClient) SearchMenu(ctx context.Context, in *SearchMenuRequest, opts ...grpc.CallOption) (*SearchMenuReply, error) {
	out := new(SearchMenuReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultClient) SearchChapter(ctx context.Context, in *SearchChapterRequest, opts ...grpc.CallOption) (*SearchChapterReply, error) {
	out := new(SearchChapterReply)
	err := c.cc.Invoke(ctx, "/novel_spider.Default/SearchChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DefaultServer is the server API for Default service.
// All implementations must embed UnimplementedDefaultServer
// for forward compatibility
type DefaultServer interface {
	// SearchNovel 搜索小说
	SearchNovel(context.Context, *SearchNovelRequest) (*SearchNovelReply, error)
	// SearchMenu 搜索目录
	SearchMenu(context.Context, *SearchMenuRequest) (*SearchMenuReply, error)
	// SearchChapter 搜索章节
	SearchChapter(context.Context, *SearchChapterRequest) (*SearchChapterReply, error)
	mustEmbedUnimplementedDefaultServer()
}

// UnimplementedDefaultServer must be embedded to have forward compatible implementations.
type UnimplementedDefaultServer struct {
}

func (UnimplementedDefaultServer) SearchNovel(context.Context, *SearchNovelRequest) (*SearchNovelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchNovel not implemented")
}
func (UnimplementedDefaultServer) SearchMenu(context.Context, *SearchMenuRequest) (*SearchMenuReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMenu not implemented")
}
func (UnimplementedDefaultServer) SearchChapter(context.Context, *SearchChapterRequest) (*SearchChapterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChapter not implemented")
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

func _Default_SearchNovel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchNovelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchNovel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchNovel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchNovel(ctx, req.(*SearchNovelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Default_SearchMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchMenu(ctx, req.(*SearchMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Default_SearchChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChapterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServer).SearchChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/novel_spider.Default/SearchChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServer).SearchChapter(ctx, req.(*SearchChapterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Default_serviceDesc = grpc.ServiceDesc{
	ServiceName: "novel_spider.Default",
	HandlerType: (*DefaultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchNovel",
			Handler:    _Default_SearchNovel_Handler,
		},
		{
			MethodName: "SearchMenu",
			Handler:    _Default_SearchMenu_Handler,
		},
		{
			MethodName: "SearchChapter",
			Handler:    _Default_SearchChapter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "novel_spider/service.proto",
}
