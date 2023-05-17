package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	storage "github.com/huanmengerkong/example-kratos/pkg/data_storage"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

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
	mysqlConnect := storage.DataBase{c.Database.Source, c.Database.Driver}
	mdb, err := storage.ConnectMysqlDb(&mysqlConnect)
	if err != nil {

	}
	redisConnect := storage.RedisBase{
		Network:      c.Redis.Network,
		Addr:         c.Redis.Addr,
		User:         "",
		Pwd:          "",
		Db:           0,
		ReadTimeout:  nil,
		WriteTimeout: nil,
	}
	rdb := storage.ConnectRedisDb(&redisConnect)
	if err != nil {

	}
	return &Data{mdb: mdb, rdb: rdb}, cleanup, nil
}
