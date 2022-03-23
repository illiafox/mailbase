package site

import (
	"mailbase/database"
	"mailbase/server/handlers/api/methods"
	"net/http"
)

func Handler(db *database.Database) http.Handler {
	rootHandler := http.NewServeMux()

	rootHandler.HandleFunc("/register", db.Wrap(methods.Reg))
	rootHandler.HandleFunc("/verify", db.Wrap(methods.Verify))
	rootHandler.HandleFunc("/login", db.Wrap(methods.Login))
	rootHandler.HandleFunc("/forgot", db.Wrap(methods.Forgot))
	rootHandler.HandleFunc("/reset", db.Wrap(methods.Reset))

	return rootHandler
}
