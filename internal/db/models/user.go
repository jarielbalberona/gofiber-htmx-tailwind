package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"null"`
	LastName  string `gorm:"null"`
}
