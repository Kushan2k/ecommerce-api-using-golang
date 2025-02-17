package user

import (
	"fmt"
	"net/http"

	"github.com/ecom-api/models"
	"github.com/ecom-api/services/auth"
	"github.com/ecom-api/types"
	"github.com/ecom-api/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


type UserService struct {
	db *gorm.DB
}



func NewUserService(database *gorm.DB) *UserService {
	return &UserService{
		db: database,
	}
}


func (s *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register",s.RegisterUser).Methods("POST")
}




func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("register user")
	
	var payload types.RegisterBodyType

	err:=utils.ParseJSON(r,payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
	}

	
	var u models.User
	results:=s.db.Where("email = ?", payload.Email).First(&u)
	// _,err=s.store.GetUserByEmail(payload.Email)

	if results.Error==nil {
		utils.WriteError(w,http.StatusInternalServerError,results.Error)
		return
	}

	if results.RowsAffected !=0{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user already exists with email %s",payload.Email))
		return
	}
	var passwordhash string=""

	passwordhash,err=auth.HashPassword(payload.Password)

	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	
	s.db.Create(&models.User{
		Email:payload.Email,
		Password:passwordhash,
		FirstName: payload.FistName,
		LastName: payload.LastName,
	})


	utils.WriteJSON(w,http.StatusCreated,fmt.Sprintf("user created with email %s",passwordhash))

	// err=s.store.CreateUser(&types.User{
	// 	Email:payload.Email,
	// 	Password:passwordhash,
	// 	FirstName: payload.FistName,
	// 	LastName: payload.LastName,
	// })

	// if err!=nil{
	// 	utils.WriteError(w,http.StatusInternalServerError,err)
	// 	return
	// }

	// utils.WriteJSON(w,http.StatusCreated,fmt.Sprintf("user created with email %s",payload.Email))


}
