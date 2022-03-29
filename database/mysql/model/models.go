package model

import "time"

//nolint:revive
type Time struct {
	Created_at time.Time `gorm:"type:DATETIME;index;not_null"`
}

//nolint:revive
type Users struct {
	User_id int `gorm:"primary_key; auto_increment; not_null" json:"-"`

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
	User_id int    `gorm:"type:INT"` // СДЕЛАЛ foreign key to Users.User_id
	Key     string `gorm:"type:VARCHAR(209);not_null"`
	Time
}
