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
	"strings"
)

// Report creates new report
func Report(db *database.Database, user *model.Users, w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	problem := r.FormValue("problem")
	if problem == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'problem' form field is empty")
		return
	}

	if len(problem) < 100 {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "Your story is too short! At least 100 bytes")
		return
	}

	problem = strings.TrimPrefix(problem, "\n")
	problem = strings.TrimSuffix(problem, "\n")

	err = db.MySQL.Reports.New(&model.Reports{
		User_id: user.User_id,
		Problem: problem,
	})

	if err != nil {
		if errors.As(err, &public.InternalWithError{}) {
			templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
			log.Println(fmt.Errorf("API: report: mysql: new report: %w", err))
		} else {
			templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
		}
		return
	}

	templates.Successful.WriteAny(w, "Thanks for your report!<br>we will send mail as soon as we fix the problem")
}
