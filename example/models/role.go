package models

import (
	"time"

	"gorm.io/gorm"
)

// gormfields:query
type Role struct {
	gorm.Model
	RoleInfo    string
	ExpiredTime time.Time
}

func (r Role) TableName() string {
	return "roles"
}
