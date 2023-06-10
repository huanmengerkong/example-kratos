package data

import (
	"context"
	"gorm.io/gorm"
	"user/internal/biz"
	"user/model"
	v1 "user/protogo/adminuser/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
	cfd  *Config
}

// NewGreeterRepo . 这个地方是实例化biz 模型对象
func NewUserRepo(data *Data, cf *Config, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		cfd:  cf,
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

func (r *UserRepo) GetInfo(ctx context.Context, request *v1.LoginRequest) (m model.FrontUser, err error) {
	err = r.data.mdb.Table(model.FrontUser{}.TableName()).Scopes(func(db *gorm.DB) *gorm.DB {
		if request.Email != "" {
			db.Where("email=? ", request.Email)
		}
		return db
	}).Where("deleted_at = 0 and status = ?", model.STATUS_USER).Find(&m).Error
	return
}

// InsertUser 新增前台注册
func (r *UserRepo) InsertUser(ctx context.Context, request model.FrontUser) (m model.FrontUser, err error) {
	err = r.data.mdb.Table(model.FrontUser{}.TableName()).Create(&request).Error
	return request, err
}
