package database

import (
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mailbase/database/mysql"
	"mailbase/database/redis"
	"mailbase/mail"
	"mailbase/shared/public"
	"mailbase/util/config"
	"net/http"
	"time"
)

type Database struct {
	Redis redis.Redis
	MySQL mysql.MySQL
	Mail  mail.Mail
}

func (db *Database) Close() (err [2]error) {
	err[0] = db.MySQL.Client.Close()
	err[1] = db.Redis.Client.Close()
	return
}

func (db *Database) Wrap(i any) http.HandlerFunc {
	switch i.(type) {

	case func(http.ResponseWriter, *http.Request):
		return i.(func(http.ResponseWriter, *http.Request))

	case func(*Database, http.ResponseWriter, *http.Request):

		f := i.(func(*Database, http.ResponseWriter, *http.Request))

		return func(writer http.ResponseWriter, request *http.Request) {
			f(db, writer, request)
		}

	default:
		panic(fmt.Sprintf("%T not implemented by wrapper", i))
	}

	return nil
}

func NewDatabase(conf config.Config) (*Database, error) {

	// // Redis
	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Pass,
		DB:       conf.Redis.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("connecting to redis: %w", err)
	}

	// // MySQL

	gormSQL, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.MySQL.Login,
			conf.MySQL.Pass,
			conf.MySQL.Protocol,
			conf.MySQL.Ip,
			conf.MySQL.Port,
			conf.MySQL.DbName,
		),
	)

	if err != nil {
		return nil, fmt.Errorf("connecting to mysql: %w", err)
	}

	sqlDB, err := mysql.Init(gormSQL)
	if err != nil {
		return nil, fmt.Errorf("initializing gorm: %w", err)
	}

	return &Database{
		// // Redis
		Redis: redis.Redis{
			Client: rdb,
			Expire: time.Duration(public.Redis.ExpireSeconds) * time.Second,
		},

		// // Mysql
		MySQL: sqlDB,

		// // Mail
		Mail: mail.NewMail(conf),
	}, nil

}
