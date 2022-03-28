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
	Client *gorm.DB
}

// Verify key: UUID from cookie, int: user_id
func (db Session) Verify(key string) (int, error) {
	session := model.Sessions{}

	err := db.Client.First(&session, "`key` = ?", key).Error // key в sql распознается как синтаксис, поэтому берем в ` `
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return -1, public.Session.NoSession
		}
		return -1, public.NewInternalWithError(err)
	}

	if int(time.Since(session.Created_at).Hours())/24 >= public.Session.SessionTimoutDays {
		err = db.Client.Delete(&session, "`key` = ?", key).Error
		if err != nil {
			log.Println(fmt.Errorf("mysql: delete old session (%s): %w", key, err))
		}
		return -1, public.Session.OldSession
	}

	return session.User_id, nil
}

// Insert creates new session
func (db Session) Insert(userid int, key string) error {
	err := db.Client.Create(&model.Sessions{
		User_id: userid,
		Key:     key,
	}).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}

	return nil
}

// Clear deletes old sessions using DATEDIFF
func (db Session) Clear(days int) error {
	err := db.Client.Delete(&model.Sessions{}, "DATEDIFF(NOW(),created_at) > ?", days).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return public.NewInternalWithError(err)
		}
	}
	return nil
}

func (db Session) DeleteByKey(key string) error {
	err := db.Client.Delete(&model.Sessions{Key: key}).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}
	return nil
}

// DeleteByUserID error can be only internal
func (db Session) DeleteByUserID(id int) error {
	return db.Client.Delete(&model.Sessions{}, "user_id = ?", id).Error
}
