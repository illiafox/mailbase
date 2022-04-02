package admin

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/server/handlers/admin/methods"
	"net/http"
)

func Handler(db *database.Database) http.Handler {
	handler := http.NewServeMux()

	// Main admin panel
	handler.Handle("/", db.Wrap(methods.Panel))

	// Admin control SUPER ADMIN ONLY
	handler.Handle("/level", db.Wrap(methods.Admins))

	// Down or up server SUPER ADMIN ONLY
	handler.Handle("/maintenance", db.Wrap(methods.Maintenance))

	// Answer the report and send mail to user
	handler.Handle("/api/check", db.Wrap(methods.Answer))

	// Get reports in json
	handler.Handle("/api/reports", db.Wrap(methods.GetReports))

	// View one report in html
	handler.Handle("/report", db.Wrap(methods.ViewReport))

	return handler
}
