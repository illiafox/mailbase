package mysql

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
)

func (db *MySQL) MailExist(email string) (*model.Users, error) {
	var user model.Users
	err := db.Client.First(&user, "email = ?", email).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, public.NewInternalWithError(err)
	}
	return &user, nil
}

func (db *MySQL) GetUserByID(id int) (*model.Users, error) {
	var user model.Users
	err := db.Client.First(&user, "user_id = ?", id).Error
	if err != nil {
		return nil, public.NewInternalWithError(err)
	}
	return &user, nil
}
