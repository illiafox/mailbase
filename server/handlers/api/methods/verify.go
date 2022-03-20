package methods

import (
	"fmt"
	"log"
	"mailbase/database"
	"mailbase/shared/public"
	"mailbase/shared/templates"
	"net/http"
)

func Verify(db *database.Database, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'key' element in query not found")
		return
	}

	user, err := db.Redis.GetBuf(key)
	if err != nil {
		if internal, ok := err.(public.InternalWithError); ok {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("API: verify: redis: get buf: %w", internal))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	exist, err := db.MySQL.MailExist(user.Email)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: verify: check login exist: %w", err))
		return
	}
	if exist != nil {
		templates.Error.WriteAny(w, public.Login.MailExist)
		return
	}

	err = db.MySQL.RegisterUser(user)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: verify: mysql: create user: %w", err))
	} else {
		templates.Successful.WriteAny(w, "Now you can <a href=\"/login\"> Login</a>")
	}
}
