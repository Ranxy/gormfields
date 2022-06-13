package example

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

func (u User) TableName() string {
	return "users"
}
