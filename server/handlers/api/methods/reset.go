package methods

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/server/handlers/api/methods/reset"
	"github.com/illiafox/mailbase/shared/templates"
	"net/http"
)

// ResetPass
// POST REQUEST: verifies form and sets new password
// GET REQUEST: finds verify event by key and returns html form
func ResetPass(db *database.Database, w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'key' element in query not found")
		return
	}
	switch r.Method {
	// GET REQUEST: finds verify event by key and returns html form
	case http.MethodGet:
		reset.Form(key, db, w)

	// POST REQUEST: verifies form and sets new password
	case http.MethodPost:
		reset.Update(key, db, w, r)
	default:
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "wrong method! only GET and POST are allowed")
	}
}

//
