package types

import "time"

type RegisterBodyType struct {
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`

}


type UserStore interface {
	GetUserByEmail(email string) (*User,error)
}


type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`

}
