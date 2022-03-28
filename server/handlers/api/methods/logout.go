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

func Logout(db *database.Database, w http.ResponseWriter, r *http.Request) {

	key, err := cookie.Session.GetClaim(r)
	if err != nil {
		if internal, ok := err.(public.InternalWithError); ok {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
			log.Println(fmt.Errorf("API: logout: cookie: get claim: %w", internal))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	err = db.MySQL.DeleteSessionByKey(key)
	if err != nil { // only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: logout: mysql: Delete Session by key (%s): %w", key, err))
		return
	}

	err = cookie.Session.DeleteClaim(w, r)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		return
	}

	templates.Successful.WriteAny(w, `You will be redirected to login page after 5 seconds
<script>  
setTimeout(function (){
    window.location = "/login";
}, 5000);
</script>  
`)

}
