package mysql

import (
	"github.com/jinzhu/gorm"
	"mailbase/database/mysql/model"
)

type MySQL struct {
	Client *gorm.DB
}

func Init(client *gorm.DB) (MySQL, error) {
	client.AutoMigrate(
		&model.Users{},
		&model.Sessions{},
	)

	return MySQL{
		Client: client,
	}, nil
}
