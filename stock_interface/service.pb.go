// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: stock/service.proto

package stock_interface

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpiredTime int64  `protobuf:"varint,2,opt,name=expiredTime,proto3" json:"expiredTime,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReply) GetExpiredTime() int64 {
	if x != nil {
		return x.ExpiredTime
	}
	return 0
}

type ListOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListOperationRequest) Reset() {
	*x = ListOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationRequest) ProtoMessage() {}

func (x *ListOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationRequest.ProtoReflect.Descriptor instead.
func (*ListOperationRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListOperationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ListOperationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockNum      string `protobuf:"bytes,1,opt,name=stockNum,proto3" json:"stockNum,omitempty"`
	Cycle         int32  `protobuf:"varint,2,opt,name=cycle,proto3" json:"cycle,omitempty"`
	OperationType int32  `protobuf:"varint,3,opt,name=operationType,proto3" json:"operationType,omitempty"`
	OperationDate int64  `protobuf:"varint,4,opt,name=operationDate,proto3" json:"operationDate,omitempty"`
	CreatedAt     int64  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     int64  `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *ListOperationReply) Reset() {
	*x = ListOperationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOperationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationReply) ProtoMessage() {}

func (x *ListOperationReply) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationReply.ProtoReflect.Descriptor instead.
func (*ListOperationReply) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListOperationReply) GetStockNum() string {
	if x != nil {
		return x.StockNum
	}
	return ""
}

func (x *ListOperationReply) GetCycle() int32 {
	if x != nil {
		return x.Cycle
	}
	return 0
}

func (x *ListOperationReply) GetOperationType() int32 {
	if x != nil {
		return x.OperationType
	}
	return 0
}

func (x *ListOperationReply) GetOperationDate() int64 {
	if x != nil {
		return x.OperationDate
	}
	return 0
}

func (x *ListOperationReply) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ListOperationReply) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpsetOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockNum      string `protobuf:"bytes,1,opt,name=stockNum,proto3" json:"stockNum,omitempty"`
	Cycle         int32  `protobuf:"varint,2,opt,name=cycle,proto3" json:"cycle,omitempty"`
	OperationType int32  `protobuf:"varint,3,opt,name=operationType,proto3" json:"operationType,omitempty"`
	Token         string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpsetOperationRequest) Reset() {
	*x = UpsetOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsetOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsetOperationRequest) ProtoMessage() {}

func (x *UpsetOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsetOperationRequest.ProtoReflect.Descriptor instead.
func (*UpsetOperationRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpsetOperationRequest) GetStockNum() string {
	if x != nil {
		return x.StockNum
	}
	return ""
}

func (x *UpsetOperationRequest) GetCycle() int32 {
	if x != nil {
		return x.Cycle
	}
	return 0
}

func (x *UpsetOperationRequest) GetOperationType() int32 {
	if x != nil {
		return x.OperationType
	}
	return 0
}

func (x *UpsetOperationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpsetOperationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsetOperationReply) Reset() {
	*x = UpsetOperationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsetOperationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsetOperationReply) ProtoMessage() {}

func (x *UpsetOperationReply) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsetOperationReply.ProtoReflect.Descriptor instead.
func (*UpsetOperationReply) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{5}
}

type DeleteOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockNum string `protobuf:"bytes,1,opt,name=stockNum,proto3" json:"stockNum,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DeleteOperationRequest) Reset() {
	*x = DeleteOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperationRequest) ProtoMessage() {}

func (x *DeleteOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperationRequest.ProtoReflect.Descriptor instead.
func (*DeleteOperationRequest) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOperationRequest) GetStockNum() string {
	if x != nil {
		return x.StockNum
	}
	return ""
}

func (x *DeleteOperationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteOperationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteOperationReply) Reset() {
	*x = DeleteOperationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stock_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOperationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOperationReply) ProtoMessage() {}

func (x *DeleteOperationReply) ProtoReflect() protoreflect.Message {
	mi := &file_stock_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOperationReply.ProtoReflect.Descriptor instead.
func (*DeleteOperationReply) Descriptor() ([]byte, []int) {
	return file_stock_service_proto_rawDescGZIP(), []int{7}
}

var File_stock_service_proto protoreflect.FileDescriptor

var file_stock_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x3e, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x44, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xce, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x73,
	0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x4a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x16, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0x9e, 0x02, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x2f, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x47, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0e, 0x55, 0x70,
	0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x7a, 0x68, 0x61, 0x6f, 0x68, 0x75, 0x61, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stock_service_proto_rawDescOnce sync.Once
	file_stock_service_proto_rawDescData = file_stock_service_proto_rawDesc
)

func file_stock_service_proto_rawDescGZIP() []byte {
	file_stock_service_proto_rawDescOnce.Do(func() {
		file_stock_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stock_service_proto_rawDescData)
	})
	return file_stock_service_proto_rawDescData
}

var file_stock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stock_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),           // 0: stock.LoginRequest
	(*LoginReply)(nil),             // 1: stock.LoginReply
	(*ListOperationRequest)(nil),   // 2: stock.ListOperationRequest
	(*ListOperationReply)(nil),     // 3: stock.ListOperationReply
	(*UpsetOperationRequest)(nil),  // 4: stock.UpsetOperationRequest
	(*UpsetOperationReply)(nil),    // 5: stock.UpsetOperationReply
	(*DeleteOperationRequest)(nil), // 6: stock.DeleteOperationRequest
	(*DeleteOperationReply)(nil),   // 7: stock.DeleteOperationReply
}
var file_stock_service_proto_depIdxs = []int32{
	0, // 0: stock.Default.Login:input_type -> stock.LoginRequest
	2, // 1: stock.Default.ListOperation:input_type -> stock.ListOperationRequest
	4, // 2: stock.Default.UpsetOperation:input_type -> stock.UpsetOperationRequest
	6, // 3: stock.Default.DeleteOperation:input_type -> stock.DeleteOperationRequest
	1, // 4: stock.Default.Login:output_type -> stock.LoginReply
	3, // 5: stock.Default.ListOperation:output_type -> stock.ListOperationReply
	5, // 6: stock.Default.UpsetOperation:output_type -> stock.UpsetOperationReply
	7, // 7: stock.Default.DeleteOperation:output_type -> stock.DeleteOperationReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stock_service_proto_init() }
func file_stock_service_proto_init() {
	if File_stock_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stock_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_stock_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_stock_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOperationRequest); i {
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
		file_stock_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOperationReply); i {
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
		file_stock_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsetOperationRequest); i {
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
		file_stock_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsetOperationReply); i {
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
		file_stock_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOperationRequest); i {
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
		file_stock_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOperationReply); i {
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
			RawDescriptor: file_stock_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stock_service_proto_goTypes,
		DependencyIndexes: file_stock_service_proto_depIdxs,
		MessageInfos:      file_stock_service_proto_msgTypes,
	}.Build()
	File_stock_service_proto = out.File
	file_stock_service_proto_rawDesc = nil
	file_stock_service_proto_goTypes = nil
	file_stock_service_proto_depIdxs = nil
}
