package mysql

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/jinzhu/gorm"
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
