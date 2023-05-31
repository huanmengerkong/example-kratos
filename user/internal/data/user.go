package data

import (
	"context"
	"user/internal/biz"
	"user/model"
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

func (r *UserRepo) GetInfo(ctx context.Context, request *v1.LoginRequest) v1.ReplyFrontedInfo {
	return v1.ReplyFrontedInfo{}
}

// InsertUser 新增前台注册
func (r *UserRepo) InsertUser(ctx context.Context, request model.FrontUser) (m model.FrontUser, err error) {
	err = r.data.mdb.Table(model.FrontUser{}.TableName()).Create(&request).Error
	return request, err
}
