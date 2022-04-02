package methods

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/templates"
	"github.com/illiafox/mailbase/util/maintenance"
	"net/http"
)

func Panel(_ *database.Database, user *model.Users, w http.ResponseWriter, r *http.Request) {
	templates.Admin.Panel.Tmpl.Execute(w, panelJSON{
		Users: user,
		Works: maintenance.Works(),
	})
}

type panelJSON struct {
	*model.Users
	Works bool
}
