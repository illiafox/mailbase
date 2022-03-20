package server

import (
	"log"
	"mailbase/database"
	apiCore "mailbase/server/handlers/api"
	siteCore "mailbase/server/handlers/site"
	"mailbase/util/config"
	"net/http"
)

func Init(conn *database.Database, conf config.Config) {

	rootHandler := http.NewServeMux()

	// Api throughout '/api/'
	rootHandler.Handle("/api/", http.StripPrefix("/api", apiCore.Handler(conn)))

	// Main site
	rootHandler.Handle("/", siteCore.Handler(conn))

	log.Println("Host started at 127.0.0.1:" + conf.Host.Port)

	err := http.ListenAndServe(":"+conf.Host.Port, rootHandler)

	log.Fatalln(err)
}
