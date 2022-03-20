package mysql

import (
	"github.com/jinzhu/gorm"
	"mailbase/database/mysql/model"
	"mailbase/shared/public"
)

func (db *MySQL) MailExist(Email string) (*model.Users, error) {
	var user model.Users
	err := db.Client.First(&user, "email = ?", Email).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, public.NewInternalWithError(err)
		}
	}
	return &user, nil
}

func (db *MySQL) GetUserById(id int) (*model.Users, error) {
	var user model.Users
	err := db.Client.First(&user, "user_id = ?", id).Error
	if err != nil {
		return nil, public.NewInternalWithError(err)
	}
	return &user, nil
}
