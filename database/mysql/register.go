package mysql

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
)

func (db *MySQL) RegisterUser(user model.Users) error {
	err := db.Client.Create(&user).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}
	return nil
}
