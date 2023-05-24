package storage

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMysqlDb(conf *DataBase) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:                    conf.Driver,
		ServerVersion:                 "",
		DSN:                           conf.Source,
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

func ConnectRedisDb(conf RedisBase) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Pwd,     // no password set
		DB:       int(conf.Db), // use default DB
	})
	return
}
