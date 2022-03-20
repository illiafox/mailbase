package methods

import (
	"fmt"
	"log"
	"mailbase/cookie"
	"mailbase/database"
	"mailbase/shared/public"
	"mailbase/shared/templates"
	"net/http"
)

func Main(db *database.Database, w http.ResponseWriter, r *http.Request) {

	key, err := cookie.GetSessionKey(r)
	if err != nil { // cannot be internal
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Session.NoSession) // overwrite error due to Cookie Error
		return
	}
	id, err := db.MySQL.VerifySession(key)
	if err != nil {
		if internal, ok := err.(public.InternalWithError); ok {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("Site: mainpage: mysql: verifysession: %w", internal))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	user, err := db.MySQL.GetUserById(id)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("Site: mainpage: mysql: getuserbyid: %w", err))
		return
	}

	templates.Main.Tmpl.Execute(w, user)
}
