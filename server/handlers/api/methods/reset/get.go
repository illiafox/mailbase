package reset

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
)

// Form finds verify event by key and returns html form
func Form(key string, db *database.Database, w http.ResponseWriter) {
	id, err := db.Redis.Forgot.Get(key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("API: reset: GET: redis: get Forgot buf: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	err = db.Redis.Reset.New(id, key)
	if err != nil { // can be only internal
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("API: reset: GET: new ResetPass buf: %w", err))
		return
	}

	templates.Reset.WriteAny(w, key)
}
