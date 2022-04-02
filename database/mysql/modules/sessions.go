package modules

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Session struct {
	Conn *gorm.DB
}

// Verify key: UUID from cookie, int: user_id
func (db Session) Verify(key string) (int, error) {
	session := model.Sessions{}

	err := db.Conn.First(&session, "`key` = ?", key).Error // key в sql распознается как синтаксис, поэтому берем в ` `
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return -1, public.Session.NoSession
		}
		return -1, public.NewInternalWithError(err)
	}

	if int(time.Since(session.Created_at).Hours())/24 >= public.Session.SessionTimoutDays {
		err = db.Conn.Delete(&session, "`key` = ?", key).Error
		if err != nil {
			log.Println(fmt.Errorf("mysql: delete old session (%s): %w", key, err))
		}
		return -1, public.Session.OldSession
	}

	return session.User_id, nil
}

// Insert creates new session
func (db Session) Insert(userid int, key string) error {
	return public.NewInternalWithError(
		db.Conn.Create(&model.Sessions{
			User_id: userid,
			Key:     key,
		}).Error,
	)
}

// Clear deletes old sessions using DATEDIFF
func (db Session) Clear(days int) error {
	err := db.Conn.Delete(&model.Sessions{}, "DATEDIFF(NOW(),created_at) > ?", days).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return public.NewInternalWithError(err)
		}
	}
	return nil
}

// DeleteByKey deletes sessions by uuid key
func (db Session) DeleteByKey(key string) error {
	return public.NewInternalWithError(
		db.Conn.Delete(&model.Sessions{Key: key}).Error,
	)
}

// DeleteByUserID deletes session by user_id value
func (db Session) DeleteByUserID(id int) error {
	return public.NewInternalWithError(
		db.Conn.Delete(&model.Sessions{}, "user_id = ?", id).Error,
	)
}
