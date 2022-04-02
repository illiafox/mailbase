package methods

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/database/mysql/modules"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
	"strconv"
)

func GetReports(db *database.Database, _ *model.Users, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, GET only", r.Method))
		return
	}

	var (
		offset, limit int
		err           error
	)
	query := r.URL.Query()
	{
		off := query.Get("offset")
		if off == "" {
			public.JSON.WriteErrorString(w, "offset query not found!")
			return
		}
		offset, err = strconv.Atoi(off)
		if err != nil {
			public.JSON.WriteErrorString(w, "cant parse offset number")
			return
		}
	}

	{
		limitStr := query.Get("limit")
		if limitStr == "" {
			public.JSON.WriteErrorString(w, "limit query not found!")
			return
		}
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			public.JSON.WriteErrorString(w, "cant parse limit number")
			return
		}
	}

	reports, err := db.MySQL.Reports.GetReports(offset, limit, query.Get("checked") == "true")
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			public.JSON.WriteError(w, public.Internal, http.StatusInternalServerError)
			log.Println(fmt.Errorf("ADMIN: API: get reports: mysql: get reports %w", err))
		} else {
			public.JSON.WriteError(w, err)
		}
		return
	}

	json.NewEncoder(w).Encode(getJSON{
		Ok:      true,
		Reports: reports,
	})
}

type getJSON struct {
	Ok      bool                     `json:"ok"`
	Reports []modules.ReportWithMail `json:"reports"`
}
