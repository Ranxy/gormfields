package models

import "gorm.io/gorm"

type StatusType uint

// gormfields:query
type User struct {
	gorm.Model
	UserName    string  `json:"user_name" gorm:"user_name"`
	UserDisplay *string `json:"user_display" gorm:"user_display"`
	Phone       int64
	Status      StatusType
}

type UserRole struct {
	gorm.Model
	UserID   uint
	RoleID   uint
	RoleDesc string
}
