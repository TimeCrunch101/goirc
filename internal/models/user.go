package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	User string
	Nick string
	Host string
}
