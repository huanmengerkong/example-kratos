package service

import (
	"context"
	"fmt"
	"user/internal/helper"
	"user/model"
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

func (s *UserService) FrontedLogin(c context.Context, req *v1.LoginRequest) (*v1.RegisterReply, error) {

	return nil, nil
}
func (s *UserService) FrontedRegister(c context.Context, req *v1.LoginRequest) (*v1.RegisterReply, error) {
	var user model.FrontUser
	users := helper.StructToGromStruct(req, user)
	if value, ok := users.(model.FrontUser); ok {
		user, err := s.uc.InsertInfo(c, value)
		return &v1.RegisterReply{
			Code: "",
			Info: &v1.ReplyFrontedInfo{
				Email:     user.Email,
				Name:      user.Name,
				CreatedAt: int64(user.CreatedAt),
				Coin:      0,
			},
			Token: "",
		}, err
	}
	return nil, nil
}

func (s *UserService) FrontedReset(c context.Context, req *v1.UserRequest) (*v1.UserRequest, error) {
	return nil, nil
}
func (s *UserService) FrontedInfo(c context.Context, req *v1.FrontedInfoRequest) (*v1.ReplyFrontedInfo, error) {
	return nil, nil

}
