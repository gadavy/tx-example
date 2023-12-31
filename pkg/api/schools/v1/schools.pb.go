// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: tx-example/api/schools/v1/schools.proto

package schools

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type School struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Addr string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *School) Reset() {
	*x = School{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *School) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*School) ProtoMessage() {}

func (x *School) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use School.ProtoReflect.Descriptor instead.
func (*School) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{0}
}

func (x *School) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *School) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *School) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *School) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Addr string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateReq) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type CreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *School `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateResp) Reset() {
	*x = CreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResp) ProtoMessage() {}

func (x *CreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResp.ProtoReflect.Descriptor instead.
func (*CreateResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResp) GetData() *School {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{3}
}

func (x *ListReq) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListReq) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*School `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{4}
}

func (x *ListResp) GetList() []*School {
	if x != nil {
		return x.List
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResp) Reset() {
	*x = DeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResp) ProtoMessage() {}

func (x *DeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_schools_v1_schools_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResp.ProtoReflect.Descriptor instead.
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_schools_v1_schools_proto_rawDescGZIP(), []int{6}
}

var File_tx_example_api_schools_v1_schools_proto protoreflect.FileDescriptor

var file_tx_example_api_schools_v1_schools_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x78, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x47, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x22, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x41, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x78, 0x5f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x0c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x32, 0xf2, 0x02,
	0x0a, 0x0a, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x41, 0x50, 0x49, 0x12, 0x78, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x74,
	0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x70, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22,
	0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x78, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x61, 0x64, 0x61, 0x76, 0x79, 0x2f, 0x74, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tx_example_api_schools_v1_schools_proto_rawDescOnce sync.Once
	file_tx_example_api_schools_v1_schools_proto_rawDescData = file_tx_example_api_schools_v1_schools_proto_rawDesc
)

func file_tx_example_api_schools_v1_schools_proto_rawDescGZIP() []byte {
	file_tx_example_api_schools_v1_schools_proto_rawDescOnce.Do(func() {
		file_tx_example_api_schools_v1_schools_proto_rawDescData = protoimpl.X.CompressGZIP(file_tx_example_api_schools_v1_schools_proto_rawDescData)
	})
	return file_tx_example_api_schools_v1_schools_proto_rawDescData
}

var file_tx_example_api_schools_v1_schools_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tx_example_api_schools_v1_schools_proto_goTypes = []interface{}{
	(*School)(nil),     // 0: tx_example.api.schools.v1.School
	(*CreateReq)(nil),  // 1: tx_example.api.schools.v1.CreateReq
	(*CreateResp)(nil), // 2: tx_example.api.schools.v1.CreateResp
	(*ListReq)(nil),    // 3: tx_example.api.schools.v1.ListReq
	(*ListResp)(nil),   // 4: tx_example.api.schools.v1.ListResp
	(*DeleteReq)(nil),  // 5: tx_example.api.schools.v1.DeleteReq
	(*DeleteResp)(nil), // 6: tx_example.api.schools.v1.DeleteResp
}
var file_tx_example_api_schools_v1_schools_proto_depIdxs = []int32{
	0, // 0: tx_example.api.schools.v1.CreateResp.data:type_name -> tx_example.api.schools.v1.School
	0, // 1: tx_example.api.schools.v1.ListResp.list:type_name -> tx_example.api.schools.v1.School
	1, // 2: tx_example.api.schools.v1.SchoolsAPI.Create:input_type -> tx_example.api.schools.v1.CreateReq
	3, // 3: tx_example.api.schools.v1.SchoolsAPI.List:input_type -> tx_example.api.schools.v1.ListReq
	5, // 4: tx_example.api.schools.v1.SchoolsAPI.Delete:input_type -> tx_example.api.schools.v1.DeleteReq
	2, // 5: tx_example.api.schools.v1.SchoolsAPI.Create:output_type -> tx_example.api.schools.v1.CreateResp
	4, // 6: tx_example.api.schools.v1.SchoolsAPI.List:output_type -> tx_example.api.schools.v1.ListResp
	6, // 7: tx_example.api.schools.v1.SchoolsAPI.Delete:output_type -> tx_example.api.schools.v1.DeleteResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tx_example_api_schools_v1_schools_proto_init() }
func file_tx_example_api_schools_v1_schools_proto_init() {
	if File_tx_example_api_schools_v1_schools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tx_example_api_schools_v1_schools_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*School); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResp); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_tx_example_api_schools_v1_schools_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResp); i {
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
			RawDescriptor: file_tx_example_api_schools_v1_schools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tx_example_api_schools_v1_schools_proto_goTypes,
		DependencyIndexes: file_tx_example_api_schools_v1_schools_proto_depIdxs,
		MessageInfos:      file_tx_example_api_schools_v1_schools_proto_msgTypes,
	}.Build()
	File_tx_example_api_schools_v1_schools_proto = out.File
	file_tx_example_api_schools_v1_schools_proto_rawDesc = nil
	file_tx_example_api_schools_v1_schools_proto_goTypes = nil
	file_tx_example_api_schools_v1_schools_proto_depIdxs = nil
}
