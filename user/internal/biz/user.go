package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/model"
	v1 "user/protogo/adminuser/v1"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1..String(), "user not found")
)

// GreeterRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *v1.UserRequest) (*v1.UserRequest, error)
	/*	Update(context.Context, *v1.UserRequest) (*v1.UserRequest, error)
		FindByID(context.Context, int64) (*v1.Admin, error)
		/*ListByHello(context.Context, string) (v1.AdminListReply, error)*/
	ListAll(context.Context) ([]v1.AdminListReply, error)

	// GetInfo fronted 前台登录
	GetInfo(ctx context.Context, request *v1.LoginRequest) v1.ReplyFrontedInfo
	// InsertUser 新增前台注册
	InsertUser(ctx context.Context, request model.FrontUser) (model.FrontUser, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, g *v1.UserRequest) (*v1.UserRequest, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v")
	return uc.repo.Save(ctx, g)
}

func (uc *UserUsecase) InsertInfo(ctx context.Context, request model.FrontUser) (model.FrontUser, error) {
	user, err := uc.repo.InsertUser(ctx, request)
	return user, err
}
