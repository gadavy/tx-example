// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: tx-example/api/auth/v1/auth.proto

package auth

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
	AuthAPI_SignIn_FullMethodName  = "/tx_example.api.users.v1.AuthAPI/SignIn"
	AuthAPI_SignUp_FullMethodName  = "/tx_example.api.users.v1.AuthAPI/SignUp"
	AuthAPI_SignOut_FullMethodName = "/tx_example.api.users.v1.AuthAPI/SignOut"
)

// AuthAPIClient is the client API for AuthAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthAPIClient interface {
	// SignIn request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | false           |
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	// SignUp request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | false           |
	SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error)
	// SignOut request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | true            |
	SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error)
}

type authAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthAPIClient(cc grpc.ClientConnInterface) AuthAPIClient {
	return &authAPIClient{cc}
}

func (c *authAPIClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, AuthAPI_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error) {
	out := new(SignUpResp)
	err := c.cc.Invoke(ctx, AuthAPI_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authAPIClient) SignOut(ctx context.Context, in *SignOutReq, opts ...grpc.CallOption) (*SignOutResp, error) {
	out := new(SignOutResp)
	err := c.cc.Invoke(ctx, AuthAPI_SignOut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthAPIServer is the server API for AuthAPI service.
// All implementations must embed UnimplementedAuthAPIServer
// for forward compatibility
type AuthAPIServer interface {
	// SignIn request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | false           |
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	// SignUp request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | false           |
	SignUp(context.Context, *SignUpReq) (*SignUpResp, error)
	// SignOut request.
	//
	// ## Authorization
	// | Parameter      | Name           |
	// | ------------- | --------------- |
	// | Session       | true            |
	SignOut(context.Context, *SignOutReq) (*SignOutResp, error)
	mustEmbedUnimplementedAuthAPIServer()
}

// UnimplementedAuthAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAuthAPIServer struct {
}

func (UnimplementedAuthAPIServer) SignIn(context.Context, *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthAPIServer) SignUp(context.Context, *SignUpReq) (*SignUpResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthAPIServer) SignOut(context.Context, *SignOutReq) (*SignOutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthAPIServer) mustEmbedUnimplementedAuthAPIServer() {}

// UnsafeAuthAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthAPIServer will
// result in compilation errors.
type UnsafeAuthAPIServer interface {
	mustEmbedUnimplementedAuthAPIServer()
}

func RegisterAuthAPIServer(s grpc.ServiceRegistrar, srv AuthAPIServer) {
	s.RegisterService(&AuthAPI_ServiceDesc, srv)
}

func _AuthAPI_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthAPI_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthAPI_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignUp(ctx, req.(*SignUpReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthAPI_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthAPIServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthAPI_SignOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthAPIServer).SignOut(ctx, req.(*SignOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthAPI_ServiceDesc is the grpc.ServiceDesc for AuthAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tx_example.api.users.v1.AuthAPI",
	HandlerType: (*AuthAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _AuthAPI_SignIn_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _AuthAPI_SignUp_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthAPI_SignOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx-example/api/auth/v1/auth.proto",
}
