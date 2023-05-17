// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: adminuser/v1/user.proto

package v1

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

// AdminUserClient is the client API for AdminUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminUserClient interface {
	// Sends a greeting
	AdminList(ctx context.Context, in *AdminListRequest, opts ...grpc.CallOption) (*AdminListReply, error)
	AdminAdd(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRequest, error)
}

type adminUserClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminUserClient(cc grpc.ClientConnInterface) AdminUserClient {
	return &adminUserClient{cc}
}

func (c *adminUserClient) AdminList(ctx context.Context, in *AdminListRequest, opts ...grpc.CallOption) (*AdminListReply, error) {
	out := new(AdminListReply)
	err := c.cc.Invoke(ctx, "/api.adminuser.v1.AdminUser/AdminList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminUserClient) AdminAdd(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRequest, error) {
	out := new(UserRequest)
	err := c.cc.Invoke(ctx, "/api.adminuser.v1.AdminUser/adminAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminUserServer is the server API for AdminUser service.
// All implementations must embed UnimplementedAdminUserServer
// for forward compatibility
type AdminUserServer interface {
	// Sends a greeting
	AdminList(context.Context, *AdminListRequest) (*AdminListReply, error)
	AdminAdd(context.Context, *UserRequest) (*UserRequest, error)
	mustEmbedUnimplementedAdminUserServer()
}

// UnimplementedAdminUserServer must be embedded to have forward compatible implementations.
type UnimplementedAdminUserServer struct {
}

func (UnimplementedAdminUserServer) AdminList(context.Context, *AdminListRequest) (*AdminListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminList not implemented")
}
func (UnimplementedAdminUserServer) AdminAdd(context.Context, *UserRequest) (*UserRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminAdd not implemented")
}
func (UnimplementedAdminUserServer) mustEmbedUnimplementedAdminUserServer() {}

// UnsafeAdminUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminUserServer will
// result in compilation errors.
type UnsafeAdminUserServer interface {
	mustEmbedUnimplementedAdminUserServer()
}

func RegisterAdminUserServer(s grpc.ServiceRegistrar, srv AdminUserServer) {
	s.RegisterService(&AdminUser_ServiceDesc, srv)
}

func _AdminUser_AdminList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminUserServer).AdminList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.adminuser.v1.AdminUser/AdminList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminUserServer).AdminList(ctx, req.(*AdminListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminUser_AdminAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminUserServer).AdminAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.adminuser.v1.AdminUser/adminAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminUserServer).AdminAdd(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminUser_ServiceDesc is the grpc.ServiceDesc for AdminUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.adminuser.v1.AdminUser",
	HandlerType: (*AdminUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminList",
			Handler:    _AdminUser_AdminList_Handler,
		},
		{
			MethodName: "adminAdd",
			Handler:    _AdminUser_AdminAdd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adminuser/v1/user.proto",
}
