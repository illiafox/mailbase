package site

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/server/handlers/site/methods"
	"github.com/illiafox/mailbase/server/middleware"
	"net/http"
)

var StaticHandler = http.FileServer(http.Dir("../shared/static"))

func Handler(db *database.Database) http.Handler {
	handler := http.NewServeMux()

	// Favicon
	handler.Handle("/favicon.ico", http.FileServer(http.Dir("../shared/images")))

	// Static
	handler.Handle("/register/", StaticHandler)
	handler.Handle("/login/", StaticHandler)
	handler.Handle("/forgot/", StaticHandler)
	handler.Handle("/report/", StaticHandler)

	middle := middleware.New(db)

	// Main page
	handler.HandleFunc("/", middle.ByLevel(model.UserLevel, methods.Main))

	return handler
}
