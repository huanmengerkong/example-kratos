package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
	"user/helper"
	"user/model"
	v1 "user/protogo/adminuser/v1"

	"user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	v1.UnimplementedAdminUserServer
	uc   *biz.UserUsecase
	h    *helper.Helper
	hjwt *helper.Hjwt
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase, jwtKey string) *UserService {
	fmt.Println("==========", jwtKey, "================")
	return &UserService{uc: uc, hjwt: helper.NewJwt(jwtKey)}
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
	data, err := s.uc.GetFrontInfo(c, req)
	if err != nil {
		return &v1.RegisterReply{}, err
	}
	token, err := s.hjwt.GetToken(c, data)
	if err != nil {
		return &v1.RegisterReply{}, err
	}
	return &v1.RegisterReply{
		Code: "",
		Info: &v1.ReplyFrontedInfo{
			Id:        data.ID,
			Email:     data.Email,
			Name:      data.Name,
			CreatedAt: int64(data.CreatedAt),
			Coin:      0,
		},
		Token: token,
	}, err
}
func (s *UserService) FrontedRegister(c context.Context, req *v1.LoginRequest) (*v1.RegisterReply, error) {
	var user model.FrontUser
	users := s.h.StructToStruct(*req, user)
	if value, ok := users.(model.FrontUser); ok {
		info, err := s.uc.InsertInfo(c, value)
		if err != nil {
			return &v1.RegisterReply{}, err
		}
		token, err := s.hjwt.GetToken(c, info)
		if err != nil {
			return &v1.RegisterReply{}, err
		}

		return &v1.RegisterReply{
			Code: "",
			Info: &v1.ReplyFrontedInfo{
				Id:        info.ID,
				Email:     info.Email,
				Name:      info.Name,
				CreatedAt: int64(user.CreatedAt),
				Coin:      0,
			},
			Token: token,
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

func (s *UserService) Server() middleware.Middleware {
	fmt.Println(111)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				auths := strings.SplitN(header.RequestHeader().Get("Authorization"), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
					return nil, errors.New("报错了")
				}
				jwtToken := auths[1]
				data, err := s.hjwt.ParamToken(ctx, jwtToken)
				fmt.Println(data, err)
				return handler(ctx, req)
			}
			return nil, errors.New("报错了")
		}
	}
}
func (s *UserService) NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.adminuser.v1.AdminUser/frontedRegister"] = struct{}{}
	whiteList["/api.adminuser.v1.AdminUser/frontedLogin"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
