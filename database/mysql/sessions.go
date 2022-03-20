package mysql

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mailbase/database/mysql/model"
	"mailbase/shared/public"
	"time"
)

// VerifySession key: UUID from cookie, int: user_id
func (db *MySQL) VerifySession(key string) (int, error) {
	session := model.Sessions{}
	hashedPass := sha256.Sum256([]byte(key))

	err := db.Client.First(&session, "`key` = ?", hex.EncodeToString(hashedPass[:])).Error // key в sql распознается как синтаксис, поэтому берем в ` `
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return -1, public.Session.NoSession
		} else {
			return -1, public.NewInternalWithError(err)
		}
	}

	if int(time.Since(session.Created_at).Hours())/24 >= public.Session.SessionTimoutDays {
		err = db.Client.Delete(&session, "key = ?", key).Error
		if err != nil {
			log.Println(fmt.Errorf("mysql: delete old session (%s): %w", key, err))
		}
		return -1, public.Session.OldSession

	}

	return session.User_id, nil
}

func (db *MySQL) InsertSession(userid int, key string) error {
	hashedPass := sha256.Sum256([]byte(key))

	err := db.Client.Create(&model.Sessions{
		User_id: userid,
		Key:     hex.EncodeToString(hashedPass[:]),
	}).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}

	return nil
}
func (db *MySQL) ClearSessions(days int) error {
	err := db.Client.Delete(&model.Sessions{}, "DATEDIFF(NOW(),created_at) > ?", days).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return public.NewInternalWithError(err)
		}
	}
	return nil
}

func (db *MySQL) DeleteSession(key string) error {
	err := db.Client.Delete(&model.Sessions{Key: key}).Error
	if err != nil {
		return public.NewInternalWithError(err)
	}
	return nil
}
