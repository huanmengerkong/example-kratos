package service

import (
	"context"
	v1 "user/protogo/adminuser/v1"

	"user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	v1.UnimplementedAdminUserServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements adminuser.GreeterServer.
func (s *UserService) AdminList(context.Context, *v1.AdminListRequest) (*v1.AdminListReply, error) {
	return &v1.AdminListReply{AdminList: []*v1.Admin{}}, nil
}
func (s *UserService) AdminAdd(context.Context, *v1.UserRequest) (*v1.UserRequest, error) {
	return &v1.UserRequest{
		Password: "",
		Name:     "",
		Status:   0,
	}, nil
}
