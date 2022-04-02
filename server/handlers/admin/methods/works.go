package methods

import (
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/util/maintenance"
	"net/http"
)

// Maintenance blocks users access with reason ('off') OR enable it ('on') via 'state' form field
func Maintenance(_ *database.Database, _ *model.Users, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, POST only", r.Method))
		return
	}

	if err := r.ParseForm(); err != nil {
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
