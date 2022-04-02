package reset

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/crypt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/validator"
	"log"
	"net/http"
)

// Update verifies form and sets new password
func Update(key string, db *database.Database, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	Password := r.FormValue("password")
	if Password == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'password' form field is empty")
		return
	}

	validator.Password(w, r, Password)
	if r.Close {
		return
	}

	id, err := db.Redis.Reset.Get(key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("API: reset: POST: redis: get ResetPass buf: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	err = db.MySQL.Session.DeleteByUserID(id)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: reset: POST: redis: DeleteSessionByUserId: %w", err))
	}

	hashedPass, err := crypt.HashPassword(Password)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: reset: hash password: %w", err))
		return
	}

	err = db.MySQL.Reset.UpdatePass(id, hashedPass)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("API: reset: POST: mysql: Update Password: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}
	templates.Successful.WriteAny(w, "Password had been updated<br> Please, <a href='/login'>Login</a>")
}
