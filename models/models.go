package models

import (
	"gorm.io/gorm"
)
type Role string

const (
	admin    Role = "admin"
	customer  Role = "customer"
	seller Role = "seller"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Auto-incrementing primary key
	Name     string `gorm:"size:100;not null"`       // Name field with max length 100
	Email    string `gorm:"unique;not null"`        // Unique email field
	Password string `gorm:"not null"`               // Password field (hashed)
	FirstName string `gorm:"size:100;not null"`       // First name field with max length 100
	LastName string `gorm:"size:100;not null"`       // Last name field with max length 100
	Role Role `gorm:"default:customer"`       // Role field with default value customer
       // Foreign Key
	Shop Shop `gorm:"foreignKey:OwnerID"`        // Foreign Key
	OTP int `gorm:"null"`       // OTP field with max length 6
	Verified bool `gorm:"default:false"`       // Verified field with default value false
}


type Shop struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Auto-incrementing primary key
	Name     string `gorm:"size:100;not null"`       // Name field with max length 100
	OwnerID    uint `gorm:"not null;unique"`        // Unique email field
	Products []Product `gorm:"foreignKey:ShopID"`        // Password field (hashed)
}

type Product struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"` // Auto-incrementing primary key
	Name     string `gorm:"size:100;not null"`       // Name field with max length 100
	Price    float64 `gorm:"not null"`        // Unique email field
	ShopID uint `gorm:"not null"`// Foreign Key        // Password field (hashed)
}
