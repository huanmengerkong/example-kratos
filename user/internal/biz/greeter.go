package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "user/protogo/adminuser/v1"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1..String(), "user not found")
)

// GreeterRepo is a Greater repo.
type UserRepo interface {
	Save(context.Context, *v1.UserRequest) (*v1.UserRequest, error)
	Update(context.Context, *v1.UserRequest) (*v1.UserRequest, error)
	FindByID(context.Context, int64) (*v1.Admin, error)
	/*ListByHello(context.Context, string) (v1.AdminListReply, error)*/
	ListAll(context.Context) ([]v1.AdminListReply, error)
}

type UserBiz struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserBiz {
	return &UserBiz{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserBiz) CreateGreeter(ctx context.Context, g *v1.UserRequest) (*v1.UserRequest, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v")
	return uc.repo.Save(ctx, g)
}
