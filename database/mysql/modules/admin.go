package modules

import (
	"fmt"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
	"net/mail"
	"strconv"
)

type Admin struct {
	Conn *gorm.DB
}

// Remove sets User.Level to 0 ONLY SUPER ADMIN
func (db Admin) Remove(identifier string) error {
	var (
		user model.Users
		err  error
	)

	switch {
	// If mail
	case isMail(identifier):
		err = db.Conn.First(&user, "email = ?", identifier).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return public.Login.MailNotFound
			}
			return public.NewInternalWithError(err)
		}

	// If id (number)
	case isNum(identifier):
		err = db.Conn.First(&user, "user_id = ?", identifier).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return public.Admin.IDNotFound
			}
			return public.NewInternalWithError(err)
		}

	default:
		return fmt.Errorf("wrong action format: %s", identifier)
	}

	switch user.Level {
	case model.UserLevel:
		return public.Admin.NotAdmin
	// cant edit super admins
	case model.SuperLevel:
		return public.Admin.EditSuper
	}

	user.Level = model.UserLevel

	return public.NewInternalWithError(
		db.Conn.Save(&user).Error,
	)
}

// Grant sets User.Level to 1 ONLY SUPER ADMIN
func (db Admin) Grant(identifier string) error {
	var (
		user model.Users
		err  error
	)
	switch {
	case isMail(identifier):
		err = db.Conn.First(&user, "email = ?", identifier).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return public.Login.MailNotFound
			}
			return public.NewInternalWithError(err)
		}

	case isNum(identifier):
		err = db.Conn.First(&user, "user_id = ?", identifier).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return public.Admin.IDNotFound
			}
			return public.NewInternalWithError(err)
		}

	default:
		return fmt.Errorf("wrong identifier format: %s", identifier)
	}

	if user.Level > 0 {
		return public.Admin.AdminAlready
	}

	user.Level = 1

	return public.NewInternalWithError(
		db.Conn.Save(&user).Error,
	)
}

func isMail(identifier string) bool {
	_, err := mail.ParseAddress(identifier)
	return err == nil
}

func isNum(identifier string) bool {
	_, err := strconv.Atoi(identifier)
	return err == nil
}
