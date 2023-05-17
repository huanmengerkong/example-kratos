package data

import (
	"context"
	v1 "user/protogo/adminuser/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) *UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, g *v1.UserRequest) (*v1.UserRequest, error) {
	return g, nil
}

func (r *UserRepo) Update(ctx context.Context, g *v1.UserRequest) (*v1.UserRequest, error) {
	return g, nil
}

func (r *UserRepo) FindByID(context.Context, int64) (*v1.Admin, error) {
	return nil, nil
}

func (r *UserRepo) ListAll(context.Context) ([]v1.AdminListReply, error) {
	return nil, nil
}
