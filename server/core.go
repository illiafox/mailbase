package server

import (
	"github.com/illiafox/mailbase/database"
	apiCore "github.com/illiafox/mailbase/server/handlers/api"
	siteCore "github.com/illiafox/mailbase/server/handlers/site"
	"github.com/illiafox/mailbase/util/config"
	"log"
	"net/http"
	"os"
)

func Init(conn *database.Database, conf config.Config, c chan os.Signal) {

	rootHandler := http.NewServeMux()

	// Api throughout '/api/'
	rootHandler.Handle("/api/", http.StripPrefix("/api", apiCore.Handler(conn)))

	// Main site
	rootHandler.Handle("/", siteCore.Handler(conn))

	log.Println("Host started at 127.0.0.1:" + conf.Host.Port)

	err := http.ListenAndServe(":"+conf.Host.Port, rootHandler)
	log.Println(err)

	c <- os.Interrupt
}
