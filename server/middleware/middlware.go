package middleware

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
)

type Middleware struct {
	*database.Database
}

type MiddleFunc func(*database.Database, *model.Users, http.ResponseWriter, *http.Request)

func (db Middleware) ByLevel(level int, handle MiddleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := cookie.Session.GetClaim(r)
		if err != nil {
			if errors.As(err, &public.InternalWithError{}) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
				log.Println(fmt.Errorf("ADMIN: panel: cookie: get claim: %w", err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		id, err := db.MySQL.Session.Verify(key)
		if err != nil {
			if errors.As(err, &public.InternalWithError{}) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
				log.Println(fmt.Errorf("ADMIN: panel: mysql: verifysession: %w", err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		user, err := db.MySQL.Login.GetUserByID(id)
		if err != nil {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: panel: mysql: getuserbyid: %w", err))
			return
		}

		if user.Level < level {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Admin.NoRights)
			return
		}

		handle(db.Database, user, w, r)
	}
}

func New(db *database.Database) Middleware {
	return Middleware{Database: db}
}
