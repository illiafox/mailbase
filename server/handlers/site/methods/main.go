package methods

import (
	"fmt"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
)

func Main(db *database.Database, w http.ResponseWriter, r *http.Request) {

	key, err := cookie.Session.GetClaim(r)
	if err != nil {
		if internal, ok := err.(public.InternalWithError); ok {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("SITE: mainpage: cookie: get claim: %w", internal))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	id, err := db.MySQL.VerifySession(key)
	if err != nil {
		if internal, ok := err.(public.InternalWithError); ok {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("SITE: mainpage: mysql: verifysession: %w", internal))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	user, err := db.MySQL.GetUserById(id)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("site: mainpage: mysql: getuserbyid: %w", err))
		return
	}

	templates.Main.Tmpl.Execute(w, user)
}
