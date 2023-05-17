package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/huanmengerkong/example-kratos/pkg/data_storage"
	"gorm.io/gorm"
	"user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	mdb *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mysqlConnect := data_storage.DataBase{c.Database.Source, c.Database.Driver}
	mdb, err := data_storage.ConnectMysqlDb(mysqlConnect)
	if err != nil {

	}
	redisConnect := data_storage.RedisBase{c.Redis.Addr, c.Redis.Addr}
	rdb, err := data_storage.redisConnect(mysqlConnect)
	if err != nil {

	}
	return &Data{mdb: mdb, rdb: rdb}, cleanup, nil
}
