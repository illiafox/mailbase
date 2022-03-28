package modules

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
)

type Register struct {
	Client *gorm.DB
}

func (db Register) NewUser(user model.Users) error {
	err := db.Client.Create(&user).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}
	return nil
}
