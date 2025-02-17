package user

import (
	"database/sql"
	"fmt"

	"github.com/ecom-api/types"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User,error){

	rows,err:=s.db.Query("SELECT id,email,first_name,last_name,password,created_at FROM users WHERE email=?",email)

	if err!=nil{
		return nil,err
	}

	u:=new(types.User)

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.Password, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		
	}
	if u.ID==0 {
		return nil,fmt.Errorf("user not found")
	}

	return u,nil



}

func (s *Store) CreateUser(u *types.User) error{
	_,err:=s.db.Exec("INSERT INTO users(email,first_name,last_name,password) VALUES(?,?,?,?)",u.Email,u.FirstName,u.LastName,u.Password)

	if err!=nil{
		return err
	}

	return nil
}
