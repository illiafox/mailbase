package methods

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

func Reset(db *database.Database, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'key' element in query not found")
		return
	}
	switch r.Method {
	// GET: Send Form
	case http.MethodGet:
		id, err := db.Redis.Forgot.Get(key)
		if err != nil {
			if errors.Is(err, public.ErrorInternal) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
				log.Println(fmt.Errorf("API: reset: GET: redis: get Forgot buf: %w", err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		err = db.Redis.Reset.New(id, key)
		if err != nil { // can be only internal
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
			log.Println(fmt.Errorf("API: reset: GET: new Reset buf: %w", err))
			return
		}

		templates.Reset.WriteAny(w, key)
	// //

	// Parse form
	case http.MethodPost:
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
			if errors.Is(err, public.ErrorInternal) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
				log.Println(fmt.Errorf("API: reset: POST: redis: get Reset buf: %w", err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		err = db.MySQL.Session.DeleteByUserID(id)
		if err != nil { // can be only internal
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
			log.Println(fmt.Errorf("API: reset: POST: redis: DeleteSessionByUserId: %w", err))
		}

		hashedPass, err := crypt.HashPassword(Password)
		if err != nil {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
			log.Println(fmt.Errorf("API: reset: hash password: %w", err))
			return
		}

		err = db.MySQL.Reset.UpdatePass(id, hashedPass)
		if err != nil {
			if errors.Is(err, public.ErrorInternal) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
				log.Println(fmt.Errorf("API: reset: POST: mysql: Update Password: %w", err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}
		templates.Successful.WriteAny(w, "Password had been updated<br> Please, <a href='/login'>Login</a>")
	}
}
