// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tx-example/api/teachers/v1/teachers.proto

package teachers

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
	TeachersAPI_Create_FullMethodName = "/tx_example.api.teachers.v1.TeachersAPI/Create"
	TeachersAPI_List_FullMethodName   = "/tx_example.api.teachers.v1.TeachersAPI/List"
	TeachersAPI_Delete_FullMethodName = "/tx_example.api.teachers.v1.TeachersAPI/Delete"
)

// TeachersAPIClient is the client API for TeachersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeachersAPIClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
}

type teachersAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTeachersAPIClient(cc grpc.ClientConnInterface) TeachersAPIClient {
	return &teachersAPIClient{cc}
}

func (c *teachersAPIClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, TeachersAPI_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersAPIClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, TeachersAPI_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teachersAPIClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, TeachersAPI_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeachersAPIServer is the server API for TeachersAPI service.
// All implementations must embed UnimplementedTeachersAPIServer
// for forward compatibility
type TeachersAPIServer interface {
	Create(context.Context, *CreateReq) (*CreateResp, error)
	List(context.Context, *ListReq) (*ListResp, error)
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
	mustEmbedUnimplementedTeachersAPIServer()
}

// UnimplementedTeachersAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTeachersAPIServer struct {
}

func (UnimplementedTeachersAPIServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTeachersAPIServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTeachersAPIServer) Delete(context.Context, *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTeachersAPIServer) mustEmbedUnimplementedTeachersAPIServer() {}

// UnsafeTeachersAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeachersAPIServer will
// result in compilation errors.
type UnsafeTeachersAPIServer interface {
	mustEmbedUnimplementedTeachersAPIServer()
}

func RegisterTeachersAPIServer(s grpc.ServiceRegistrar, srv TeachersAPIServer) {
	s.RegisterService(&TeachersAPI_ServiceDesc, srv)
}

func _TeachersAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersAPI_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersAPIServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersAPI_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersAPIServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeachersAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeachersAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeachersAPI_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeachersAPIServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TeachersAPI_ServiceDesc is the grpc.ServiceDesc for TeachersAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeachersAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tx_example.api.teachers.v1.TeachersAPI",
	HandlerType: (*TeachersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TeachersAPI_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TeachersAPI_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TeachersAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx-example/api/teachers/v1/teachers.proto",
}