package service

import (
	"context"
	"fmt"
	v1 "helloworld/api/helloworld/v1"
)

type UseSrv struct {
	v1.UnimplementedUserServer
}

func NewUserSrv() *UseSrv {
	return &UseSrv{}
}
func (receiver *UseSrv) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	return nil, nil
}

func (receiver *UseSrv) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	return nil, nil
}
func (receiver *UseSrv) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	return nil, nil
}
func (receiver *UseSrv) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	fmt.Println("user into")
	return &v1.GetUserReply{User: &v1.CreateUserRequest{
		Name:  "xx",
		Passd: "",
		Email: "11@qq.com",
	}}, nil
}
func (receiver *UseSrv) ListUser(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	return nil, nil
}
