package modules

import (
	"github.com/illiafox/mailbase/database/mysql/model"
	"github.com/illiafox/mailbase/shared/public"
	"github.com/jinzhu/gorm"
	"time"
)

type Reports struct {
	Conn *gorm.DB
}

// New creates report from struct
func (db Reports) New(report *model.Reports) error {
	err := db.Conn.First(report, "checked = false AND user_id = ?", report.User_id).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return public.NewInternalWithError(err)
		}
	} else {
		return public.Reports.StillProcessing
	}
	return public.NewInternalWithError(
		db.Conn.Omit("checked_at").Create(report).Error,
	)
}

// Check set answer for report
func (db Reports) Check(reportID, adminID int, answer string) error {
	var report model.Reports
	err := db.Conn.First(&report, "report_id = ?", reportID).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return public.Reports.NotFound
		}
		return public.NewInternalWithError(err)
	}

	if report.Checked {
		return public.Reports.IsChecked
	}

	report.Checked = true
	report.Answer = answer
	report.Admin_id = adminID
	report.Checked_at = time.Now()

	return public.NewInternalWithError(
		db.Conn.Save(&report).Error,
	)
}

type ReportWithMail struct {
	Email   string
	Created string
	model.Reports
}

// GetReports returns slice with reports, nil if not exists
// offset: offset relative to query; limit: how much return (max 100)
func (db Reports) GetReports(offset, limit int, checked bool) ([]ReportWithMail, error) {
	if limit > 100 {
		limit = 100
	}

	var reports []ReportWithMail

	scan := db.Conn.Model(model.Reports{}).
		Select(`report_id,user_id,problem,answer,checked,
		(SELECT email FROM users WHERE reports.user_id = users.user_id) AS email,
		DATE_FORMAT(created_at, '%H:%i %d-%m-%Y') AS created`).
		Where("checked = ?", checked).
		Order("created_at DESC").
		Offset(offset).Limit(limit).Scan(&reports)
	if scan.Error != nil {
		return nil, public.NewInternalWithError(scan.Error)
	}
	if scan.RowsAffected == 0 {
		return nil, public.Admin.ReportsNotFound
	}

	return reports, nil
}

type ReportTemplate struct {
	model.Reports
	Created     string
	Checked     string
	User, Admin string
}

// Get returns report with mails and dates
func (db Reports) Get(id int) (*ReportTemplate, error) {
	var report ReportTemplate

	err := db.Conn.Model(model.Reports{}).
		Select(`*,
		(SELECT email FROM users WHERE users.user_id = reports.user_id) AS user,
		(SELECT email FROM users WHERE users.user_id = reports.admin_id) AS admin,
		DATE_FORMAT(created_at, '%H:%i %d-%m-%Y') AS created,
		DATE_FORMAT(checked_at, '%H:%i %d-%m-%Y') AS checked`).
		Where("report_id = ?", id).Scan(&report).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, public.Reports.NotFound
		}
		return nil, public.NewInternalWithError(err)
	}

	return &report, nil
}

// DeleteByID deletes report by report_id value
func (db Reports) DeleteByID(id int) error {
	return public.NewInternalWithError(db.Conn.Delete(&model.Reports{Report_id: id}).Error)
}
