package mysql

import (
	"github.com/jinzhu/gorm"
	"mailbase/database/mysql/model"
	"mailbase/shared/public"
)

func (db *MySQL) ResetPass(id int, pass string) error {
	var user model.Users

	err := db.Client.First(&user, "user_id = ?", id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return public.Login.MailNotFound
		} else {
			return public.NewInternalWithError(err)
		}
	}

	if user.Password == pass {
		return public.Forgot.SamePassword
	}

	return db.Client.Model(&user).Update("password", pass).Error
}
