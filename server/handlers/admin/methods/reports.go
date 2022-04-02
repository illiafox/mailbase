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
	"strconv"
)

// ViewReport write report in html
func ViewReport(db *database.Database, user *model.Users, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, GET only", r.Method))
		return
	}

	var (
		reportID int
		err      error
	)
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
			log.Println(fmt.Errorf("ADMIN: view post: mysql: get report (%d): %w", user.User_id, err))
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
