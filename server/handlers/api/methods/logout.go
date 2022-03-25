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

func Logout(db *database.Database, w http.ResponseWriter, r *http.Request) {
	key, err := cookie.GetSessionKey(r)
	if err != nil { // cannot be internal
		templates.Error.WriteAnyCode(w, http.StatusForbidden, err) // overwrite error due to Cookie Error
		return
	}

	err = db.MySQL.DeleteSessionByKey(key)
	if err != nil { // only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.InternalError)
		log.Println(fmt.Errorf("API: logout: mysql: Delete Session by key (%s): %w", key, err))
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
