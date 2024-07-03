package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Name     struct {
		FirstName string
		LastName  string
	} `gorm:"embedded"`
	Role    Role `gorm:"type:role;default:'USER'"`
	College string
}
