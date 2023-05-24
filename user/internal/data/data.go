package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	storage "github.com/huanmengerkong/example-kratos/pkg/data_storage"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"user/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

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
	mysqlConnect := storage.DataBase{Source: c.Database.Source, Driver: c.Database.Driver}
	mdb, err := storage.ConnectMysqlDb(&mysqlConnect)
	if err != nil {
		panic(fmt.Sprintf("mdb 报错了%v", err))
	}
	redisConnect := storage.RedisBase{
		Network:      c.Redis.Network,
		Addr:         c.Redis.Addr,
		User:         c.Redis.User,
		Pwd:          c.Redis.Pwd,
		Db:           c.Redis.Db,
		ReadTimeout:  c.Redis.ReadTimeout,
		WriteTimeout: c.Redis.WriteTimeout,
	}
	rdb := storage.ConnectRedisDb(redisConnect)
	if err != nil {
		panic(fmt.Sprintf("rdb 报错了%v", err))
	}
	return &Data{mdb: mdb, rdb: rdb}, cleanup, nil
}