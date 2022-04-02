package methods

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/cookie"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/database/mysql/modules"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
	"strconv"
)

func GetReports(db *database.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		public.JSON.WriteErrorString(w, "Method not allowed! Use GET", http.StatusMethodNotAllowed)
		return
	}

	key, err := cookie.Session.GetClaim(r)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			public.JSON.WriteError(w, public.Internal, http.StatusInternalServerError)
			log.Println(fmt.Errorf("ADMIN: API: get reports: get claim: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	id, err := db.MySQL.Session.Verify(key)
	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			public.JSON.WriteError(w, public.Internal, http.StatusInternalServerError)
			log.Println(fmt.Errorf("ADMIN: API: get reports: mysql: verifysession: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	user, err := db.MySQL.Login.GetUserByID(id)
	if err != nil {
		public.JSON.WriteError(w, public.Internal, http.StatusInternalServerError)
		log.Println(fmt.Errorf("ADMIN: API: get reports: mysql: GetUserByID(%d): %w", id, err))
		return
	}

	// Master admin check
	if user.Level < model.AdminLevel {
		public.JSON.WriteError(w, public.Admin.NoRights)
		return
	}

	//

	var offset, limit int
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
