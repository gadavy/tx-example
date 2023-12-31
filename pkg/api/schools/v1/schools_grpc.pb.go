// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tx-example/api/schools/v1/schools.proto

package schools

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
	SchoolsAPI_Create_FullMethodName = "/tx_example.api.schools.v1.SchoolsAPI/Create"
	SchoolsAPI_List_FullMethodName   = "/tx_example.api.schools.v1.SchoolsAPI/List"
	SchoolsAPI_Delete_FullMethodName = "/tx_example.api.schools.v1.SchoolsAPI/Delete"
)

// SchoolsAPIClient is the client API for SchoolsAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchoolsAPIClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
}

type schoolsAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSchoolsAPIClient(cc grpc.ClientConnInterface) SchoolsAPIClient {
	return &schoolsAPIClient{cc}
}

func (c *schoolsAPIClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, SchoolsAPI_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolsAPIClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, SchoolsAPI_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schoolsAPIClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, SchoolsAPI_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchoolsAPIServer is the server API for SchoolsAPI service.
// All implementations must embed UnimplementedSchoolsAPIServer
// for forward compatibility
type SchoolsAPIServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
	mustEmbedUnimplementedSchoolsAPIServer()
}

// UnimplementedSchoolsAPIServer must be embedded to have forward compatible implementations.
type UnimplementedSchoolsAPIServer struct {
}

func (UnimplementedSchoolsAPIServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSchoolsAPIServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSchoolsAPIServer) Delete(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSchoolsAPIServer) mustEmbedUnimplementedSchoolsAPIServer() {}

// UnsafeSchoolsAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchoolsAPIServer will
// result in compilation errors.
type UnsafeSchoolsAPIServer interface {
	mustEmbedUnimplementedSchoolsAPIServer()
}

func RegisterSchoolsAPIServer(s grpc.ServiceRegistrar, srv SchoolsAPIServer) {
	s.RegisterService(&SchoolsAPI_ServiceDesc, srv)
}

func _SchoolsAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolsAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchoolsAPI_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolsAPIServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchoolsAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolsAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchoolsAPI_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolsAPIServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchoolsAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchoolsAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchoolsAPI_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchoolsAPIServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SchoolsAPI_ServiceDesc is the grpc.ServiceDesc for SchoolsAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchoolsAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tx_example.api.schools.v1.SchoolsAPI",
	HandlerType: (*SchoolsAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SchoolsAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SchoolsAPI_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SchoolsAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx-example/api/schools/v1/schools.proto",
}
