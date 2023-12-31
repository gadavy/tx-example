// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: tx-example/api/auth/v1/auth.proto

package auth

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

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignInReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignUpResp) Reset() {
	*x = SignUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResp) ProtoMessage() {}

func (x *SignUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResp.ProtoReflect.Descriptor instead.
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

type SignOutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutReq) Reset() {
	*x = SignOutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutReq) ProtoMessage() {}

func (x *SignOutReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutReq.ProtoReflect.Descriptor instead.
func (*SignOutReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

type SignOutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignOutResp) Reset() {
	*x = SignOutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignOutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResp) ProtoMessage() {}

func (x *SignOutResp) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResp.ProtoReflect.Descriptor instead.
func (*SignOutResp) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

type ScopesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScopesReq) Reset() {
	*x = ScopesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScopesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopesReq) ProtoMessage() {}

func (x *ScopesReq) ProtoReflect() protoreflect.Message {
	mi := &file_tx_example_api_auth_v1_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopesReq.ProtoReflect.Descriptor instead.
func (*ScopesReq) Descriptor() ([]byte, []int) {
	return file_tx_example_api_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

var File_tx_example_api_auth_v1_auth_proto protoreflect.FileDescriptor

var file_tx_example_api_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x09, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x3d, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x32, 0xe9,
	0x02, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x72, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x12, 0x22, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x72,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x22, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x74,
	0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x75, 0x70, 0x12, 0x76, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x23, 0x2e,
	0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x24, 0x2e, 0x74, 0x78, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x64, 0x61, 0x76, 0x79, 0x2f,
	0x74, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tx_example_api_auth_v1_auth_proto_rawDescOnce sync.Once
	file_tx_example_api_auth_v1_auth_proto_rawDescData = file_tx_example_api_auth_v1_auth_proto_rawDesc
)

func file_tx_example_api_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_tx_example_api_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_tx_example_api_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_tx_example_api_auth_v1_auth_proto_rawDescData)
	})
	return file_tx_example_api_auth_v1_auth_proto_rawDescData
}

var file_tx_example_api_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tx_example_api_auth_v1_auth_proto_goTypes = []interface{}{
	(*SignInReq)(nil),   // 0: tx_example.api.users.v1.SignInReq
	(*SignInResp)(nil),  // 1: tx_example.api.users.v1.SignInResp
	(*SignUpReq)(nil),   // 2: tx_example.api.users.v1.SignUpReq
	(*SignUpResp)(nil),  // 3: tx_example.api.users.v1.SignUpResp
	(*SignOutReq)(nil),  // 4: tx_example.api.users.v1.SignOutReq
	(*SignOutResp)(nil), // 5: tx_example.api.users.v1.SignOutResp
	(*ScopesReq)(nil),   // 6: tx_example.api.users.v1.ScopesReq
}
var file_tx_example_api_auth_v1_auth_proto_depIdxs = []int32{
	0, // 0: tx_example.api.users.v1.AuthAPI.SignIn:input_type -> tx_example.api.users.v1.SignInReq
	2, // 1: tx_example.api.users.v1.AuthAPI.SignUp:input_type -> tx_example.api.users.v1.SignUpReq
	4, // 2: tx_example.api.users.v1.AuthAPI.SignOut:input_type -> tx_example.api.users.v1.SignOutReq
	1, // 3: tx_example.api.users.v1.AuthAPI.SignIn:output_type -> tx_example.api.users.v1.SignInResp
	3, // 4: tx_example.api.users.v1.AuthAPI.SignUp:output_type -> tx_example.api.users.v1.SignUpResp
	5, // 5: tx_example.api.users.v1.AuthAPI.SignOut:output_type -> tx_example.api.users.v1.SignOutResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tx_example_api_auth_v1_auth_proto_init() }
func file_tx_example_api_auth_v1_auth_proto_init() {
	if File_tx_example_api_auth_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tx_example_api_auth_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResp); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutReq); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignOutResp); i {
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
		file_tx_example_api_auth_v1_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScopesReq); i {
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
			RawDescriptor: file_tx_example_api_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tx_example_api_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_tx_example_api_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_tx_example_api_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_tx_example_api_auth_v1_auth_proto = out.File
	file_tx_example_api_auth_v1_auth_proto_rawDesc = nil
	file_tx_example_api_auth_v1_auth_proto_goTypes = nil
	file_tx_example_api_auth_v1_auth_proto_depIdxs = nil
}
