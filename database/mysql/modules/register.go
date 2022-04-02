package modules

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
)

type Register struct {
	Conn *gorm.DB
}

// NewUser creates new user in table
func (db Register) NewUser(user model.Users) error {
	return public.NewInternalWithError(
		db.Conn.Create(&user).Error,
	)
}
