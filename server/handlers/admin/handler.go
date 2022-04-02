package admin

import (
	"github.com/illiafox/mailbase/database"
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/server/handlers/admin/methods"
	"github.com/illiafox/mailbase/server/middleware"
	"net/http"
)

func Handler(db *database.Database) http.Handler {
	handler := http.NewServeMux()

	middle := middleware.New(db)

	// Main admin panel
	handler.Handle("/", middle.ByLevel(model.AdminLevel, methods.Panel))

	// Admin control SUPER ADMIN ONLY
	handler.Handle("/level", middle.ByLevel(model.SuperLevel, methods.Admins))

	// Down or up server SUPER ADMIN ONLY
	handler.Handle("/maintenance", middle.ByLevel(model.SuperLevel, methods.Maintenance))

	// Answer the report and send mail to user
	handler.Handle("/api/check", middle.ByLevel(model.AdminLevel, methods.Answer))

	// Get reports in json
	handler.Handle("/api/reports", middle.ByLevel(model.AdminLevel, methods.GetReports))

	// View one report in html
	handler.Handle("/report", middle.ByLevel(model.AdminLevel, methods.ViewReport))

	return handler
}
