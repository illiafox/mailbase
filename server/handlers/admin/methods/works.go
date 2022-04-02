package methods

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/util/maintenance"
	"log"
	"net/http"
)

// Maintenance blocks users access with reason ('off') OR enable it ('on') via 'state' form field
func Maintenance(db *database.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		public.WriteWithCode(w, http.StatusMethodNotAllowed, "Method not allowed! Use POST")
		return
	}

	key, err := cookie.Session.GetClaim(r)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: maintenance: cookie: get claim: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	id, err := db.MySQL.Session.Verify(key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: maintenance: mysql: verifysession: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	user, err := db.MySQL.Login.GetUserByID(id)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("ADMIN: maintenance: mysql: GetUserByID(%d): %w", id, err))
		return
	}

	// Master admin check
	if user.Level < model.AdminLevel {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Admin.NoRights)
		return
	}

	if user.Level < model.SuperLevel {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Admin.NotSuper)
		return
	}

	// //

	err = r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	state := r.FormValue("state")
	if state == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'state' form field is empty")
		return
	}

	switch state {
	case "off":
		if !maintenance.Works() {
			templates.Error.WriteAny(w, "server is already down")
			return
		}
		details := r.FormValue("details")
		if details == "" {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, "'details' form field is empty")
			return
		}

		maintenance.Off(details)
		templates.Successful.WriteAny(w, "<strong>SERVER IS DOWN</strong><br>admin panel still works")
	case "on":
		if maintenance.Works() {
			templates.Error.WriteAny(w, "server is already up")
			return
		}
		maintenance.On()
		templates.Successful.WriteAny(w, "<strong>SERVER IS UP NOW</strong>")
	default:
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("wrong state format: %s", state))
	}
}
