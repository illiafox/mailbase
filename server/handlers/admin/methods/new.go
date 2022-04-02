package methods

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
)

// Admins sets User.Admins to 1 ('grant') OR 0 ('remove') ONLY SUPER ADMIN
func Admins(db *database.Database, _ *model.Users, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, POST only", r.Method))
		return
	}

	err := r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	identifier := r.FormValue("identifier")
	if identifier == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'identifier' form field is empty")
		return
	}

	action := r.FormValue("action")
	if action == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'action' form field is empty")
		return
	}

	switch action {
	case "grant":
		err = db.MySQL.Admin.Grant(identifier)
		if err != nil {
			if errors.As(err, &public.InternalWithError{}) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
				log.Println(fmt.Errorf("ADMIN: admins: mysql: grant admin (%s): %w", identifier, err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}
		templates.Successful.WriteAny(w, `User is admin now<br>Redirecting back to admin panel..
		<meta http-equiv="refresh" content="2 url=/admin">`)

	case "remove":
		err = db.MySQL.Admin.Remove(identifier)
		if err != nil {
			if errors.As(err, &public.InternalWithError{}) {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
				log.Println(fmt.Errorf("ADMIN: admins: mysql: remove admin (%s): %w", identifier, err))
			} else {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
			}
			return
		}

		templates.Successful.WriteAny(w, `User is not admin anymore<br>Redirecting back to admin panel..
		<meta http-equiv="refresh" content="2 url=/admin">`)

	default:
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("wrong action format: %s", action))
	}
}
