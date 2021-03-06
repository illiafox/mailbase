package methods

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"github.com/illiafox/mailbase/crypt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/validator"
	"log"
	"net/http"
	"net/mail"
)

// Register creates verify event to create account
func Register(db *database.Database, w http.ResponseWriter, r *http.Request) {
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
		log.Println(fmt.Errorf("API: register: check login exist: %w", err))
		return
	}
	if exist != nil {
		templates.Error.WriteAny(w, public.Login.MailExist)
		return
	}

	key := uuid.NewString()
	hash := sha256.Sum256([]byte(key))
	key = hex.EncodeToString(hash[:])

	var buf bytes.Buffer

	err = templates.Mail.Verify.WriteBytes(&buf, key)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: register: create message with key: %w", err))
		return
	}

	err = db.Mail.SendMessage(Mail, "Your verify link", buf.String())
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: register: send mail (%s): %w", Mail, err))
		return
	}

	hashedPass, err := crypt.HashPassword(Password)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: register: hash password: %w", err))
		return
	}

	err = db.Redis.Verify.New(model.Users{
		Email:    Mail,
		Password: hashedPass,
	}, key)

	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: register: new buf: %w", err))
	} else {
		templates.Successful.WriteAny(w, "Check your email box :)")
	}
}
