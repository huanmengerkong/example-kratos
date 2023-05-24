package service

import (
	"context"
	"fmt"
	v1 "user/protogo/adminuser/v1"

	"user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	v1.UnimplementedAdminUserServer
	uc *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// SayHello implements adminuser.GreeterServer.
func (s *UserService) AdminList(c context.Context, req *v1.AdminListRequest) (*v1.AdminListReply, error) {
	return &v1.AdminListReply{AdminList: []*v1.Admin{}}, nil
}
func (s *UserService) AdminAdd(c context.Context, req *v1.UserRequest) (*v1.UserRequest, error) {
	v, err := s.uc.CreateUser(c, req)
	fmt.Println(v, err)
	return &v1.UserRequest{
		Password: "",
		Name:     "",
		Status:   0,
	}, nil
}
