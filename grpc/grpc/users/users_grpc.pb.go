// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: users.proto

package users

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
	UsersAPI_GetAllUsers_FullMethodName = "/users.UsersAPI/getAllUsers"
	UsersAPI_GetUserById_FullMethodName = "/users.UsersAPI/getUserById"
	UsersAPI_CreateUser_FullMethodName  = "/users.UsersAPI/createUser"
	UsersAPI_DeleteUser_FullMethodName  = "/users.UsersAPI/deleteUser"
	UsersAPI_UpdateUser_FullMethodName  = "/users.UsersAPI/updateUser"
)

// UsersAPIClient is the client API for UsersAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersAPIClient interface {
	GetAllUsers(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetUserById(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error)
	CreateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error)
	DeleteUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error)
}

type usersAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersAPIClient(cc grpc.ClientConnInterface) UsersAPIClient {
	return &usersAPIClient{cc}
}

func (c *usersAPIClient) GetAllUsers(ctx context.Context, in *APIRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UsersAPI_GetAllUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAPIClient) GetUserById(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UsersAPI_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAPIClient) CreateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UsersAPI_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAPIClient) DeleteUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UsersAPI_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersAPIClient) UpdateUser(ctx context.Context, in *UserArgs, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UsersAPI_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersAPIServer is the server API for UsersAPI service.
// All implementations must embed UnimplementedUsersAPIServer
// for forward compatibility
type UsersAPIServer interface {
	GetAllUsers(context.Context, *APIRequest) (*UserReply, error)
	GetUserById(context.Context, *UserArgs) (*UserReply, error)
	CreateUser(context.Context, *UserArgs) (*UserReply, error)
	DeleteUser(context.Context, *UserArgs) (*UserReply, error)
	UpdateUser(context.Context, *UserArgs) (*UserReply, error)
}

// UnimplementedUsersAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUsersAPIServer struct {
}

func (UnimplementedUsersAPIServer) GetAllUsers(context.Context, *APIRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUsersAPIServer) GetUserById(context.Context, *UserArgs) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUsersAPIServer) CreateUser(context.Context, *UserArgs) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUsersAPIServer) DeleteUser(context.Context, *UserArgs) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUsersAPIServer) UpdateUser(context.Context, *UserArgs) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUsersAPIServer) mustEmbedUnimplementedUsersAPIServer() {}

// UnsafeUsersAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersAPIServer will
// result in compilation errors.
type UnsafeUsersAPIServer interface {
	mustEmbedUnimplementedUsersAPIServer()
}

func RegisterUsersAPIServer(s grpc.ServiceRegistrar, srv UsersAPIServer) {
	s.RegisterService(&UsersAPI_ServiceDesc, srv)
}

func _UsersAPI_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_GetAllUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).GetAllUsers(ctx, req.(*APIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAPI_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).GetUserById(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAPI_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).CreateUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAPI_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).DeleteUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersAPI_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersAPIServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersAPI_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersAPIServer).UpdateUser(ctx, req.(*UserArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersAPI_ServiceDesc is the grpc.ServiceDesc for UsersAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UsersAPI",
	HandlerType: (*UsersAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getAllUsers",
			Handler:    _UsersAPI_GetAllUsers_Handler,
		},
		{
			MethodName: "getUserById",
			Handler:    _UsersAPI_GetUserById_Handler,
		},
		{
			MethodName: "createUser",
			Handler:    _UsersAPI_CreateUser_Handler,
		},
		{
			MethodName: "deleteUser",
			Handler:    _UsersAPI_DeleteUser_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _UsersAPI_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
