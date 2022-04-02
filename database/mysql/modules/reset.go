package modules

import (
	"github.com/illiafox/mailbase/crypt"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
)

type Reset struct {
	Conn *gorm.DB
}

// UpdatePass sets new password for user id
func (db Reset) UpdatePass(id int, pass string) error {
	var user model.Users

	err := db.Conn.First(&user, "user_id = ?", id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return public.Login.MailNotFound
		}
		return public.NewInternalWithError(err)
	}

	if crypt.ComparePassword(user.Password, pass) {
		return public.Forgot.SamePassword
	}

	return public.NewInternalWithError(
		db.Conn.Model(&user).Update("password", pass).Error,
	)
}
