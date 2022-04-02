package methods

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/templates"
	"net/http"
)

func Main(db *database.Database, user *model.Users, w http.ResponseWriter, r *http.Request) {
	templates.Main.Tmpl.Execute(w, user)
}
