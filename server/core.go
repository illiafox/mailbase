package server

import (
	"context"
	"github.com/illiafox/mailbase/database"
	apiCore "github.com/illiafox/mailbase/server/handlers/api"
	siteCore "github.com/illiafox/mailbase/server/handlers/site"
	"github.com/illiafox/mailbase/util/config"
	"net/http"
)

type Shutdown func(context.Context) error

func Init(conn *database.Database, conf config.Config) *http.Server {
	rootHandler := http.NewServeMux()

	// Api throughout '/api/'
	rootHandler.Handle("/api/", http.StripPrefix("/api", apiCore.Handler(conn)))

	// Main site
	rootHandler.Handle("/", siteCore.Handler(conn))

	return &http.Server{
		Addr:    ":" + conf.Host.Port,
		Handler: rootHandler,
	}
}
