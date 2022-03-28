package mysql

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/database/mysql/modules"
	"github.com/jinzhu/gorm"
)

type MySQL struct {
	Client   *gorm.DB
	Login    modules.Login
	Register modules.Register
	Reset    modules.Reset
	Session  modules.Session
}

func Init(client *gorm.DB) (MySQL, error) {
	client.AutoMigrate(
		&model.Users{},
		&model.Sessions{},
	)

	return MySQL{
		Client: client,

		Login: modules.Login{Client: client},

		Register: modules.Register{Client: client},

		Reset: modules.Reset{Client: client},

		Session: modules.Session{Client: client},
	}, nil
}
