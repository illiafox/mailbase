package methods

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mailbase/database"
	"mailbase/database/mysql/model"
	"mailbase/shared/public"
	"mailbase/shared/templates"
	"net/http"
	"net/mail"
	"unicode"
)

func Reg(db *database.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		public.WriteWithCode(w, http.StatusMethodNotAllowed, "Method not allowed! Use POST")
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

	// Mail check
	_, err = mail.ParseAddress(Mail)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "Wrong mail format!\nExample: fortstudio8@gmail.com")
		return
	}

	// Password check Why not regexp? Because re2 does not support lookaheads '?= '
	count, low, up, num := 0, false, false, false
	for _, s := range Password {
		if !unicode.IsLetter(s) && !unicode.IsNumber(s) {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, "Invalid password format: Only numbers/letters are allowed")
			return
		}
		switch {
		case unicode.IsLower(s):
			low = true
		case unicode.IsUpper(s):
			up = true
		case unicode.IsNumber(s):
			num = true
		}
		count++
	}

	if public.Register.PasswordMin > count || count > public.Register.PasswordMax {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Register.InvalidLength)
		return
	}

	if !(low && up && num) {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Register.InvalidFormat)
		return
	}

	exist, err := db.MySQL.MailExist(Mail)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: register: check login exist: %w", err))
		return
	}
	if exist != nil {
		templates.Error.WriteAny(w, public.Login.MailExist)
		return
	}

	hashedPass := sha256.Sum256([]byte(Password))
	key := uuid.NewString()

	var buf bytes.Buffer
	err = templates.Mail.WriteBytes(&buf, key)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: register: create message with key: %w", err))
		return
	}

	err = db.Mail.SendMessage(Mail, "Your verify link", buf.String())
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: register: send mail (%s): %w", Mail, err))
		return
	}

	err = db.Redis.NewBuf(model.Users{
		Email:    Mail,
		Password: hex.EncodeToString(hashedPass[:]),
	}, key)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: register: new buf: %w", err))
	} else {
		templates.Successful.WriteAny(w, "Check your email box :)")
	}

}
