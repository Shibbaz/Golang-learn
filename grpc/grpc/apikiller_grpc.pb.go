// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	GetAllUsers(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*APIReply, error)
	GetUserById(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error)
	CreateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error)
	DeleteUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error)
	UpdateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error)
	CreatePost(ctx context.Context, in *PostArgs, opts ...grpc.CallOption) (*APIReply, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetAllUsers(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/getAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetUserById(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/getUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/createUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/deleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) UpdateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreatePost(ctx context.Context, in *PostArgs, opts ...grpc.CallOption) (*APIReply, error) {
	out := new(APIReply)
	err := c.cc.Invoke(ctx, "/apikiller.API/createPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	GetAllUsers(context.Context, *APIRequest) (*APIReply, error)
	GetUserById(context.Context, *UserArgs) (*APIReply, error)
	CreateUser(context.Context, *UserArgs) (*APIReply, error)
	DeleteUser(context.Context, *UserArgs) (*APIReply, error)
	UpdateUser(context.Context, *UserArgs) (*APIReply, error)
	CreatePost(context.Context, *PostArgs) (*APIReply, error)
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) GetAllUsers(context.Context, *APIRequest) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedAPIServer) GetUserById(context.Context, *UserArgs) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedAPIServer) CreateUser(context.Context, *UserArgs) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAPIServer) DeleteUser(context.Context, *UserArgs) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAPIServer) UpdateUser(context.Context, *UserArgs) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedAPIServer) CreatePost(context.Context, *PostArgs) (*APIReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/getAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAllUsers(ctx, req.(*APIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/getUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetUserById(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/createUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/deleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/updateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).UpdateUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apikiller.API/createPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreatePost(ctx, req.(*PostArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apikiller.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAllUsers",
			Handler:    _API_GetAllUsers_Handler,
		},
		{
			MethodName: "getUserById",
			Handler:    _API_GetUserById_Handler,
		},
		{
			MethodName: "createUser",
			Handler:    _API_CreateUser_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _API_DeleteUser_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _API_UpdateUser_Handler,
		},
		{
			MethodName: "createPost",
			Handler:    _API_CreatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/apikiller.proto",
}