package methods

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/crypt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/validator"
	"log"
	"net/http"
	"net/mail"
)

// Login verifies your account and creates cookies
func Login(db *database.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, POST only", r.Method))
		return
	}

	err := r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	Mail := r.FormValue("mail")
	if Mail == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'mail' form field is empty")
		return
	}

	Password := r.FormValue("password")
	if Password == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'password' form field is empty")
		return
	}

	// MailVerify check
	_, err = mail.ParseAddress(Mail)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "Wrong mail format!\nExample: fortstudio8@gmail.com")
		return
	}

	validator.Password(w, r, Password)
	if r.Close {
		return
	}

	exist, err := db.MySQL.Login.MailExist(Mail)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: login: check login exist: %w", err))
		return
	}
	if exist == nil {
		_ = templates.Error.WriteAny(w, "Account not found,<br>but you can <a href=\"/register\"> create one</a>")
		return
	}

	if !crypt.ComparePassword(exist.Password, Password) {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Login.IncorrectPassword)
		return
	}

	key := uuid.NewString()
	hashedPass := sha256.Sum256([]byte(key))
	key = hex.EncodeToString(hashedPass[:])

	_, err = cookie.Session.SetClaim(w, r, key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("API: login: cookie: set claim: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	err = db.MySQL.Session.Insert(exist.User_id, key)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: login: insert session: %w", err))
	} else {
		templates.Successful.WriteAny(w, "You can visit <a href=\"/\">main page</a> now")
	}
}
