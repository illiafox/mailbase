package api

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/server/handlers/api/methods"
	"github.com/illiafox/mailbase/server/middleware"
	"net/http"
)

func Handler(db *database.Database) http.Handler {
	handler := http.NewServeMux()

	// Create account via form
	handler.HandleFunc("/register", db.Wrap(methods.Register))

	// Verify account with key from mail
	handler.HandleFunc("/verify", db.Wrap(methods.Verify))

	// Login and create cookie session
	handler.HandleFunc("/login", db.Wrap(methods.Login))

	// Delete current session
	handler.HandleFunc("/logout", db.Wrap(methods.Logout))

	// Send mail to reset password
	handler.HandleFunc("/forgot", db.Wrap(methods.Forgot))

	// Reset password
	// GET: return form
	// POST: update pass
	handler.HandleFunc("/reset", db.Wrap(methods.ResetPass))

	middle := middleware.New(db)

	// Send report
	handler.HandleFunc("/report", middle.ByLevel(model.UserLevel, methods.Report))

	return handler
}
