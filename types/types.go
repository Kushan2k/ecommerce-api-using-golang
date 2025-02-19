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
