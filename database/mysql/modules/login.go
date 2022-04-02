package modules

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
)

type Login struct {
	Conn *gorm.DB
}

// MailExist returns user struct or nil if not exists
func (db Login) MailExist(email string) (*model.Users, error) {
	var user model.Users
	err := db.Conn.First(&user, "email = ?", email).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, public.NewInternalWithError(err)
	}
	return &user, nil
}

// GetUserByID returns user struct by id
func (db Login) GetUserByID(id int) (*model.Users, error) {
	var user model.Users

	return &user, public.NewInternalWithError(
		db.Conn.First(&user, "user_id = ?", id).Error,
	)
}
