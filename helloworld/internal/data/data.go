package data

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	mdb *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	db, err := NewMdb(c)
	if err != nil {
		panic(err)
	}
	rdb := NewRdb(c)
	return &Data{mdb: db, rdb: rdb}, cleanup, nil
}

func NewMdb(c *conf.Data) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: c.Database.Driver,
		//		ServerVersion:                 "",
		DSN:                           c.Database.Source,
		DSNConfig:                     nil,
		Conn:                          nil,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             0,
		DefaultDatetimePrecision:      nil,
		DisableWithReturning:          true,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
		DontSupportForShareClause:     true,
		DontSupportNullAsDefaultValue: false,
		DontSupportRenameColumnUnique: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return
}

func NewRdb(c *conf.Data) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Pwd,     // no password set
		DB:       int(c.Redis.Db), // use default DB
	})
	return
}
