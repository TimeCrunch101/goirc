package models

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	Name  string
	Users map[*User]bool
}
