package methods

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"mailbase/database"
	"mailbase/shared/public"
	"mailbase/shared/templates"
	"mailbase/validator"
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
		id, err := db.Redis.GetForgotPass(key)
		if err != nil {
			if internal, ok := err.(public.InternalWithError); ok {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
				log.Println(fmt.Errorf("API: reset: GET: redis: get Forgot buf: %w", internal))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		err = db.Redis.NewResetPass(id, key)
		if err != nil { // can be only internal
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
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

		id, err := db.Redis.GetResetPass(key)
		if err != nil {
			if internal, ok := err.(public.InternalWithError); ok {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
				log.Println(fmt.Errorf("API: reset: POST: redis: get Reset buf: %w", internal))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		err = db.MySQL.DeleteSessionByUserId(id)
		if err != nil { // can be only internal
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("API: reset: POST: redis: DeleteSessionByUserId: %w", err))
		}

		hashedPass := sha256.Sum256([]byte(Password))

		err = db.MySQL.ResetPass(id, hex.EncodeToString(hashedPass[:]))
		if err != nil {
			if internal, ok := err.(public.InternalWithError); ok {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
				log.Println(fmt.Errorf("API: reset: POST: mysql: Update Password: %w", internal))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		templates.Successful.WriteAny(w, "Password had been updated<br> Please, <a href='/login'>Login</a>")
	}

}
