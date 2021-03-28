// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.0
// source: novel_spider/service.proto

package novel_spider_interface

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SearchNovelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// novelName 书名
	NovelName string `protobuf:"bytes,1,opt,name=novelName,proto3" json:"novelName,omitempty"`
}

func (x *SearchNovelRequest) Reset() {
	*x = SearchNovelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchNovelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNovelRequest) ProtoMessage() {}

func (x *SearchNovelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNovelRequest.ProtoReflect.Descriptor instead.
func (*SearchNovelRequest) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchNovelRequest) GetNovelName() string {
	if x != nil {
		return x.NovelName
	}
	return ""
}

type SearchNovelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// novels 小说列表
	Novels []*NovelInfo `protobuf:"bytes,1,rep,name=novels,proto3" json:"novels,omitempty"`
}

func (x *SearchNovelReply) Reset() {
	*x = SearchNovelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchNovelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchNovelReply) ProtoMessage() {}

func (x *SearchNovelReply) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchNovelReply.ProtoReflect.Descriptor instead.
func (*SearchNovelReply) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchNovelReply) GetNovels() []*NovelInfo {
	if x != nil {
		return x.Novels
	}
	return nil
}

type SearchMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Novel *NovelInfo `protobuf:"bytes,1,opt,name=novel,proto3" json:"novel,omitempty"`
}

func (x *SearchMenuRequest) Reset() {
	*x = SearchMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMenuRequest) ProtoMessage() {}

func (x *SearchMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMenuRequest.ProtoReflect.Descriptor instead.
func (*SearchMenuRequest) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{2}
}

func (x *SearchMenuRequest) GetNovel() *NovelInfo {
	if x != nil {
		return x.Novel
	}
	return nil
}

type SearchMenuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapters []*NovelChapter `protobuf:"bytes,1,rep,name=chapters,proto3" json:"chapters,omitempty"`
	Novel    *NovelInfo      `protobuf:"bytes,2,opt,name=novel,proto3" json:"novel,omitempty"`
}

func (x *SearchMenuReply) Reset() {
	*x = SearchMenuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMenuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMenuReply) ProtoMessage() {}

func (x *SearchMenuReply) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMenuReply.ProtoReflect.Descriptor instead.
func (*SearchMenuReply) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{3}
}

func (x *SearchMenuReply) GetChapters() []*NovelChapter {
	if x != nil {
		return x.Chapters
	}
	return nil
}

func (x *SearchMenuReply) GetNovel() *NovelInfo {
	if x != nil {
		return x.Novel
	}
	return nil
}

type SearchChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter *NovelChapter `protobuf:"bytes,1,opt,name=chapter,proto3" json:"chapter,omitempty"`
	Novel   *NovelInfo    `protobuf:"bytes,2,opt,name=novel,proto3" json:"novel,omitempty"`
}

func (x *SearchChapterRequest) Reset() {
	*x = SearchChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchChapterRequest) ProtoMessage() {}

func (x *SearchChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchChapterRequest.ProtoReflect.Descriptor instead.
func (*SearchChapterRequest) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{4}
}

func (x *SearchChapterRequest) GetChapter() *NovelChapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

func (x *SearchChapterRequest) GetNovel() *NovelInfo {
	if x != nil {
		return x.Novel
	}
	return nil
}

type SearchChapterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter *NovelChapter `protobuf:"bytes,1,opt,name=chapter,proto3" json:"chapter,omitempty"`
}

func (x *SearchChapterReply) Reset() {
	*x = SearchChapterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_novel_spider_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchChapterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchChapterReply) ProtoMessage() {}

func (x *SearchChapterReply) ProtoReflect() protoreflect.Message {
	mi := &file_novel_spider_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchChapterReply.ProtoReflect.Descriptor instead.
func (*SearchChapterReply) Descriptor() ([]byte, []int) {
	return file_novel_spider_service_proto_rawDescGZIP(), []int{5}
}

func (x *SearchChapterReply) GetChapter() *NovelChapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

var File_novel_spider_service_proto protoreflect.FileDescriptor

var file_novel_spider_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f,
	0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x18, 0x6e, 0x6f, 0x76, 0x65,
	0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f,
	0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f,
	0x76, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x6f, 0x76, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x06,
	0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e,
	0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x76, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x73, 0x22, 0x42, 0x0a,
	0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x76, 0x65,
	0x6c, 0x22, 0x78, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73,
	0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x52, 0x08, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x05,
	0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f,
	0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x22, 0x7b, 0x0a, 0x14, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x6f, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c,
	0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x22, 0x4a, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x76, 0x65, 0x6c, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x32, 0xff, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x4f, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x12,
	0x20, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x4c, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6e, 0x75, 0x12,
	0x1f, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x55, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x7a, 0x68, 0x61, 0x6f, 0x68, 0x75, 0x61, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_novel_spider_service_proto_rawDescOnce sync.Once
	file_novel_spider_service_proto_rawDescData = file_novel_spider_service_proto_rawDesc
)

func file_novel_spider_service_proto_rawDescGZIP() []byte {
	file_novel_spider_service_proto_rawDescOnce.Do(func() {
		file_novel_spider_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_novel_spider_service_proto_rawDescData)
	})
	return file_novel_spider_service_proto_rawDescData
}

var file_novel_spider_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_novel_spider_service_proto_goTypes = []interface{}{
	(*SearchNovelRequest)(nil),   // 0: novel_spider.SearchNovelRequest
	(*SearchNovelReply)(nil),     // 1: novel_spider.SearchNovelReply
	(*SearchMenuRequest)(nil),    // 2: novel_spider.SearchMenuRequest
	(*SearchMenuReply)(nil),      // 3: novel_spider.SearchMenuReply
	(*SearchChapterRequest)(nil), // 4: novel_spider.SearchChapterRequest
	(*SearchChapterReply)(nil),   // 5: novel_spider.SearchChapterReply
	(*NovelInfo)(nil),            // 6: novel_spider.NovelInfo
	(*NovelChapter)(nil),         // 7: novel_spider.NovelChapter
}
var file_novel_spider_service_proto_depIdxs = []int32{
	6,  // 0: novel_spider.SearchNovelReply.novels:type_name -> novel_spider.NovelInfo
	6,  // 1: novel_spider.SearchMenuRequest.novel:type_name -> novel_spider.NovelInfo
	7,  // 2: novel_spider.SearchMenuReply.chapters:type_name -> novel_spider.NovelChapter
	6,  // 3: novel_spider.SearchMenuReply.novel:type_name -> novel_spider.NovelInfo
	7,  // 4: novel_spider.SearchChapterRequest.chapter:type_name -> novel_spider.NovelChapter
	6,  // 5: novel_spider.SearchChapterRequest.novel:type_name -> novel_spider.NovelInfo
	7,  // 6: novel_spider.SearchChapterReply.chapter:type_name -> novel_spider.NovelChapter
	0,  // 7: novel_spider.Default.SearchNovel:input_type -> novel_spider.SearchNovelRequest
	2,  // 8: novel_spider.Default.SearchMenu:input_type -> novel_spider.SearchMenuRequest
	4,  // 9: novel_spider.Default.SearchChapter:input_type -> novel_spider.SearchChapterRequest
	1,  // 10: novel_spider.Default.SearchNovel:output_type -> novel_spider.SearchNovelReply
	3,  // 11: novel_spider.Default.SearchMenu:output_type -> novel_spider.SearchMenuReply
	5,  // 12: novel_spider.Default.SearchChapter:output_type -> novel_spider.SearchChapterReply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_novel_spider_service_proto_init() }
func file_novel_spider_service_proto_init() {
	if File_novel_spider_service_proto != nil {
		return
	}
	file_novel_spider_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_novel_spider_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchNovelRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novel_spider_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchNovelReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novel_spider_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMenuRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novel_spider_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMenuReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novel_spider_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchChapterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_novel_spider_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchChapterReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_novel_spider_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_novel_spider_service_proto_goTypes,
		DependencyIndexes: file_novel_spider_service_proto_depIdxs,
		MessageInfos:      file_novel_spider_service_proto_msgTypes,
	}.Build()
	File_novel_spider_service_proto = out.File
	file_novel_spider_service_proto_rawDesc = nil
	file_novel_spider_service_proto_goTypes = nil
	file_novel_spider_service_proto_depIdxs = nil
}
