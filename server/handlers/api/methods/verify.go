package methods

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
)

func Verify(db *database.Database, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'key' element in query not found")
		return
	}

	user, err := db.Redis.Verify.Get(key)
	if err != nil {
		if errors.Is(err, public.ErrorInternal) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
			log.Println(fmt.Errorf("API: verify: redis: get buf: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	exist, err := db.MySQL.Login.MailExist(user.Email)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
		log.Println(fmt.Errorf("API: verify: check login exist: %w", err))
		return
	}
	if exist != nil {
		templates.Error.WriteAny(w, public.Login.MailExist)
		return
	}

	err = db.MySQL.Register.NewUser(user)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.ErrorInternal)
		log.Println(fmt.Errorf("API: verify: mysql: create user: %w", err))
	} else {
		templates.Successful.WriteAny(w, "Now you can <a href=\"/login\"> Login</a>")
	}
}
