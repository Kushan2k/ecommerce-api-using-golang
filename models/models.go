package models

import (
	"gorm.io/gorm"
)

// User Model (Multiple Addresses Supported)
type User struct {
	gorm.Model
	FirstName string `gorm:"size:100"`
	LastName  string `gorm:"size:100"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Phone     string `gorm:"size:20"`
	Role      string `gorm:"type:enum('customer', 'admin', 'vendor');default:'customer'"`
	Addresses []UserAddress `gorm:"foreignKey:UserID"` // One-to-Many Relationship

	OTP int `gorm:"null"`       // OTP field with max length 6
	Verified bool `gorm:"default:false"`       // Verified field with default value false
}

// User Address Model (Multiple Addresses per User)
type UserAddress struct {
	gorm.Model
	UserID      uint
	Address     string `gorm:"not null"`
	City        string `gorm:"size:100;not null"`
	State       string `gorm:"size:100"`
	PostalCode  string `gorm:"size:20"`
	Country     string `gorm:"size:100;not null"`
	IsDefault   bool   `gorm:"default:false"` // Flag for Default Address
	User        User   `gorm:"foreignKey:UserID"`
}

// Category Model
type Category struct {
	gorm.Model
	Name     string  `gorm:"unique;not null"`
	ParentID *uint   // Nullable Parent Category
	Parent   *Category `gorm:"foreignKey:ParentID"`
}

// Product Model
type Product struct {
	gorm.Model
	Name        string  `gorm:"not null"`
	Description string
	BasePrice   float64 `gorm:"not null"`
	CategoryID  uint
	VendorID    *uint
	StockQty    int `gorm:"default:0"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Vendor      *User    `gorm:"foreignKey:VendorID"`

	ProductImages []ProductImage `gorm:"foreignKey:ProductID"` // One-to-Many Relationship
}

type ProductImage struct {
	gorm.Model
	ProductID uint
	ImageURL  string `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`

}

// Product Variation Model
type ProductVariation struct {
	gorm.Model
	ProductID uint
	SKU       string  `gorm:"unique;not null"`
	Price     float64 `gorm:"not null"`
	StockQty  int     `gorm:"default:0"`
	Product   Product `gorm:"foreignKey:ProductID"`
	VariantImages []VariantImage `gorm:"foreignKey:VariantID"` // One-to-Many Relationship
}

type VariantImage struct {
	gorm.Model
	VariantID uint
	ImageURL  string `gorm:"not null"`
	Variant   ProductVariation `gorm:"foreignKey:VariantID"`
	
}

// Variation Attributes (e.g., Size, Color)
type VariationAttribute struct {
	gorm.Model
	VariationID   uint
	AttributeName string `gorm:"size:100;not null"`
	AttributeValue string `gorm:"size:100;not null"`
	Variation    ProductVariation `gorm:"foreignKey:VariationID"`
}

// Order Model
type Order struct {
	gorm.Model
	UserID      uint
	AddressID   uint  // Reference to UserAddress
	TotalPrice  float64 `gorm:"not null"`
	Status      string  `gorm:"type:enum('pending', 'shipped', 'delivered', 'cancelled', 'returned');default:'pending'"`
	User        User      `gorm:"foreignKey:UserID"`
	ShippingAddress UserAddress `gorm:"foreignKey:AddressID"`
}

// Order Items (Products in an Order)
type OrderItem struct {
	gorm.Model
	OrderID     uint
	VariationID uint
	Quantity    int     `gorm:"not null;check:quantity > 0"`
	Price       float64 `gorm:"not null"`
	Order       Order   `gorm:"foreignKey:OrderID"`
	Variation   ProductVariation `gorm:"foreignKey:VariationID"`
}

// Payment Model
type Payment struct {
	gorm.Model
	OrderID       uint
	PaymentMethod string `gorm:"type:enum('credit_card', 'paypal', 'bank_transfer', 'cash_on_delivery');not null"`
	PaymentStatus string `gorm:"type:enum('pending', 'completed', 'failed');default:'pending'"`
	TransactionID *string `gorm:"unique"`
	Order         Order   `gorm:"foreignKey:OrderID"`
}

// Shipping Model
type Shipping struct {
	gorm.Model
	OrderID         uint
	TrackingNumber  *string `gorm:"unique"`
	ShippingStatus  string  `gorm:"type:enum('pending', 'shipped', 'delivered');default:'pending'"`
	Order           Order   `gorm:"foreignKey:OrderID"`
}

// Review Model
type Review struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Rating    int    `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment   string
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

// Wishlist Model
type Wishlist struct {
	gorm.Model
	UserID    uint
	ProductID uint
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

// Cart Model
type Cart struct {
	gorm.Model
	UserID      uint
	VariationID uint
	Quantity    int `gorm:"not null;check:quantity > 0"`
	User        User             `gorm:"foreignKey:UserID"`
	Variation   ProductVariation `gorm:"foreignKey:VariationID"`
}
