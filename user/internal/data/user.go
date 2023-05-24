package data

import (
	"context"
	"user/internal/biz"
	v1 "user/protogo/adminuser/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo . 这个地方是实例化biz 模型对象
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) Save(ctx context.Context, g *v1.UserRequest) (*v1.UserRequest, error) {
	// s, _ := r.data.mdb.Scopes((ctx, &v1.UserRequest{Name: "aaa"})
	// fmt.Println(s)
	return g, nil
}

func (r *UserRepo) ListAll(context.Context) ([]v1.AdminListReply, error) {
	return nil, nil
}
