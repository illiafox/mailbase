package methods

import (
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
	"strconv"
)

// ViewReport write report in html
func ViewReport(db *database.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		public.WriteWithCode(w, http.StatusMethodNotAllowed, "Method not allowed! Use GET")
		return
	}

	key, err := cookie.Session.GetClaim(r)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: view report: cookie: get claim: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	id, err := db.MySQL.Session.Verify(key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: view report: mysql: verifysession: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	user, err := db.MySQL.Login.GetUserByID(id)
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
		log.Println(fmt.Errorf("ADMIN: view report: mysql: GetUserByID(%d): %w", id, err))
		return
	}

	// Master admin check
	if user.Level < model.AdminLevel {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, public.Admin.NoRights)
		return
	}

	var reportID int
	{
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			public.JSON.WriteErrorString(w, "id query not found!")
			return
		}
		reportID, err = strconv.Atoi(idStr)
		if err != nil {
			public.JSON.WriteErrorString(w, "cant parse id number")
			return
		}
	}

	report, err := db.MySQL.Reports.Get(reportID)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("ADMIN: view post: mysql: get report (%d): %w", id, err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}
	err = templates.Admin.Report.Tmpl.Execute(w, report)
	if err != nil {
		log.Println(err)
	}
}
