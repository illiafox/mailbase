package mysql

import "mailbase/database/mysql/model"

func (db *MySQL) ResetPass(id int, pass string) error {
	// TODO: if pass = newpass -> successful error
	return db.Client.Model(&model.Users{}).Where("user_id = ?", id).Update("password", pass).Error
}
