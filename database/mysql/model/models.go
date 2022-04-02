package model

import "time"

const (
	UserLevel = iota
	AdminLevel
	SuperLevel
)

//nolint:revive
type Time struct {
	Created_at time.Time `gorm:"type:DATETIME;index;not_null"`
}

//nolint:revive
type Users struct {
	User_id int `gorm:"primary_key; auto_increment; not_null" json:"-"`

	// Permission level
	// 0: UserLevel
	// 1: AdminLevel
	// 2: SuperLevel: can change permissions
	Level int `gorm:"default: 0; not_null"`

	// Email limit is a maximum of 64 characters (octets) in the "local part" (before the "@")
	// and a maximum of 255 characters (octets)
	// in the domain part (after the "@") for a total length of 320 characters
	Email string `gorm:"type:VARCHAR(320);not_null"`

	// Hash size is 64 bits
	Password string `gorm:"type:VARCHAR(64);not_null"`

	// Create time
	Time `json:"-"`
}

//nolint:revive
type Sessions struct {
	User_id int    `gorm:"type:INT"`
	Key     string `gorm:"type:VARCHAR(209);not_null"`
	Time
}

//nolint:revive
type Reports struct {
	Report_id int    `gorm:"primary_key; auto_increment; not_null"`
	User_id   int    `gorm:"type:INT"`
	Admin_id  int    `gorm:"type:INT"`
	Checked   bool   `gorm:"type:BOOLEAN"`
	Problem   string `gorm:"type:TEXT"`
	Answer    string `gorm:"type:TEXT"`

	Checked_at time.Time `gorm:"type:DATETIME;index;"`

	Time
}
