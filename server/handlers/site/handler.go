package site

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/server/handlers/site/methods"
	"net/http"
)

var StaticHandler = http.FileServer(http.Dir("../shared/static"))

func Handler(db *database.Database) http.Handler {
	rootHandler := http.NewServeMux()

	// Static
	rootHandler.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("../shared/images/"))))
	rootHandler.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("../shared/js/"))))
	rootHandler.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("../shared/css/"))))

	rootHandler.Handle("/register/", StaticHandler)
	rootHandler.Handle("/login/", StaticHandler)
	rootHandler.Handle("/forgot/", StaticHandler)

	rootHandler.HandleFunc("/", db.Wrap(methods.Main))

	return rootHandler
}
