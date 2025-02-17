package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Auto-incrementing primary key
	Name     string `gorm:"size:100;not null"`       // Name field with max length 100
	Email    string `gorm:"unique;not null"`        // Unique email field
	Password string `gorm:"not null"`               // Password field (hashed)
	FirstName string `gorm:"size:100;not null"`       // First name field with max length 100
	LastName string `gorm:"size:100;not null"`       // Last name field with max length 100
}
