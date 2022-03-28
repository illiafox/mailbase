package database

import (
	"context"
	"fmt"
	goRedis "github.com/go-redis/redis/v8"
	"github.com/illiafox/mailbase/database/mysql"
	"github.com/illiafox/mailbase/database/redis"
	"github.com/illiafox/mailbase/mail"
	"github.com/illiafox/mailbase/util/config"
	"github.com/jinzhu/gorm"
	//nolint:revive
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
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

func (db *Database) Wrap(f func(*Database, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(db, w, r)
	}
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
			conf.MySQL.IP,
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

	// // Mail
	Mail, err := mail.NewMail(conf)
	if err != nil {
		return nil, fmt.Errorf("initializing smtp: %w", err)
	}

	return &Database{
		// // Redis
		Redis: redis.NewRedis(rdb, conf),
		// // Mysql
		MySQL: sqlDB,
		// // Mail
		Mail: Mail,
	}, nil
}

type Wrap interface {
	func(*Database, http.ResponseWriter, *http.Request) | func(http.ResponseWriter, *http.Request)
}
