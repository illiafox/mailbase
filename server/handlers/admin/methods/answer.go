package methods

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/illiafox/mailbase/shared/templates"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Answer report ('answer') or delete it ('delete')
func Answer(db *database.Database, user *model.Users, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("method %s not supported, POST only", r.Method))
		return
	}

	err := r.ParseForm()
	if err != nil {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("form parsing error: %w", err))
		return
	}

	action := r.FormValue("action")
	if action == "" {
		templates.Error.WriteAnyCode(w, http.StatusForbidden, "'action' form field is empty")
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

		switch action {
		case "delete":
			err = db.MySQL.Reports.DeleteByID(reportID)
			if err != nil {
				if errors.As(err, &public.InternalWithError{}) {
					templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
					log.Println(fmt.Errorf("ADMIN: check: mysql: delete report by ID (%d): %w", reportID, err))
				} else {
					templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
				}
				return
			}
			templates.Successful.WriteAny(w, `Report is deleted<br>Redirecting back to admin panel..
		<meta http-equiv="refresh" content="2 url=/admin">`)

		case "answer":

			answer := r.FormValue("answer")
			if answer == "" {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, "'answer' form field is empty")
				return
			}

			if len(answer) < 100 {
				templates.Error.WriteAnyCode(w, http.StatusForbidden, "Your answer is too short! At least 100 bytes")
				return
			}

			err = db.MySQL.Reports.Check(reportID, user.User_id, answer)
			if err != nil {
				if errors.As(err, &public.InternalWithError{}) {
					templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
					log.Println(fmt.Errorf("ADMIN: check: mysql: check report by ID (%d): %w", reportID, err))
				} else {
					templates.Error.WriteAnyCode(w, http.StatusForbidden, err)
				}
				return
			}

			// Send mail

			var buf bytes.Buffer

			answer = strings.ReplaceAll(answer, "\n", "<br>")

			templates.Mail.Answer.WriteBytes(&buf, answer)

			err = db.Mail.SendMessage(user.Email, "Answer for report", buf.String())
			if err != nil {
				templates.Error.WriteAnyCode(w, http.StatusInternalServerError, public.Internal)
				log.Println(fmt.Errorf("ADMIN: answer: mail: send answer message (%s): %w", user.Email, err))
			}

			templates.Successful.WriteAny(w, `Report is checked<br>Redirecting back to admin panel..
			<meta http-equiv="refresh" content="2 url=/admin">`)

		default:
			templates.Error.WriteAnyCode(w, http.StatusForbidden, fmt.Errorf("wrong action format: %s", action))
		}
	}
}
