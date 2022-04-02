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
	Admin    modules.Admin
	Reports  modules.Reports
}

func Init(client *gorm.DB) (MySQL, error) {
	client.AutoMigrate(
		&model.Users{},
		&model.Sessions{},
		&model.Reports{},
	)

	return MySQL{
		Client: client,

		Login: modules.Login{Conn: client},

		Register: modules.Register{Conn: client},

		Reset: modules.Reset{Conn: client},

		Session: modules.Session{Conn: client},

		Admin: modules.Admin{Conn: client},

		Reports: modules.Reports{Conn: client},
	}, nil
}
