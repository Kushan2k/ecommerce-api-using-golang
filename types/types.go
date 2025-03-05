package types

import "time"

type RegisterBodyType struct {
	FirstName string `json:"first_name" validate:"max=50,required"`
	LastName string `json:"last_name" validate:"max=50,required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,required"`

}

type LoginBodyType struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8,required"`
}

type VerifyBodyType struct {
	Email string `json:"email" validate:"required,email"`
	OTP int `json:"otp" validate:"required"`
}

type ResendBodyType struct {
	Email string `json:"email" validate:"required,email"`
}


type UserStore interface {
	GetUserByEmail(email string) (*User,error)
	CreateUser(*User) error
}


type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`

}

type ShopRequest struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	LogoURL string `json:"logo_url" validate:"url"`
	UserID int `json:"user_id" validate:"required"`
}

type ProductCreateResponse struct {
	Name 			string  `json:"name" validate:"required,min=5,max=100"`
	Description 	string  `json:"description " validate:"required,min=50,max=1000"`
	BasePrice 		float64 `json:"base_price" validate:"required,gte=0"`
	CategoryID 		uint    `json:"category_id" validate:"required"`
	ShopID    		uint    `json:"shop_id" validate:"required"`
	StockQty    	int     `json:"stock_qty" validate:"required,gte=0"`
}
