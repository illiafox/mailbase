package server

import (
	"context"
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/server/handlers/admin"
	"github.com/illiafox/mailbase/server/handlers/api"
	"github.com/illiafox/mailbase/server/handlers/site"
	"github.com/illiafox/mailbase/util/config"
	"github.com/illiafox/mailbase/util/maintenance"
	"net/http"
)

type Shutdown func(context.Context) error

func Init(conn *database.Database, conf config.Config) *http.Server {
	handler := http.NewServeMux()

	// Static
	handler.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("../shared/images/"))))
	handler.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../shared/js/"))))
	handler.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("../shared/css/"))))

	// Api throughout '/api/'
	handler.Handle("/api/", http.StripPrefix("/api",
		maintenance.Handler{
			Root: api.Handler(conn),
		},
	))

	// Admin throughout '/admin/'
	handler.Handle("/admin/", http.StripPrefix("/admin", admin.Handler(conn)))

	// Main site
	handler.Handle("/",
		maintenance.Handler{
			Root: site.Handler(conn),
		},
	)

	return &http.Server{
		Addr:    ":" + conf.Host.Port,
		Handler: handler,
	}
}
