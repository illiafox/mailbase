package methods

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
	"net/mail"
)

func Forgot(db *database.Database, w http.ResponseWriter, r *http.Request) {
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

	// MailVerify check
	_, err = mail.ParseAddress(Mail)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "Wrong mail format!\nExample: fortstudio8@gmail.com")
		return
	}

	exist, err := db.MySQL.MailExist(Mail)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: forgot: check login exist: %w", err))
		return
	}
	if exist == nil {
		templates.Error.WriteAny(w, "Account not found, but you can <a href=\"/register\"> create new one</a>")
		return
	}

	key := uuid.NewString()

	var buf bytes.Buffer
	err = templates.Mail.ResetPass.WriteBytes(&buf, key)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: forgot: create message with key: %w", err))
		return
	}

	err = db.Mail.SendMessage(Mail, "Your verify link", buf.String())
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: forgot: send mail (%s): %w", Mail, err))
		return
	}

	db.Redis.Forgot.New(exist.User_id, key)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: register: new buf: %w", err))
	} else {
		templates.Successful.WriteAny(w, "Check your email box :)")
	}

}
