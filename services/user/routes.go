package user

import (
	"fmt"
	"net/http"

	"github.com/ecom-api/services/auth"
	"github.com/ecom-api/types"
	"github.com/ecom-api/utils"
	"github.com/gorilla/mux"
)


type UserService struct {
	store types.UserStore
}



func NewUserService(userStore types.UserStore) *UserService {
	return &UserService{
		store: userStore,
	}
}


func (s *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register",s.RegisterUser).Methods("POST")
}




func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	var payload types.RegisterBodyType

	err:=utils.ParseJSON(r,payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
	}

	

	_,err=s.store.GetUserByEmail(payload.Email)

	if err==nil{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("user already exists with email %s",payload.Email))
		return
	}
	var passwordhash string=""

	passwordhash,err=auth.HashPassword(payload.Password)

	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	err=s.store.CreateUser(&types.User{
		Email:payload.Email,
		Password:passwordhash,
		FirstName: payload.FistName,
		LastName: payload.LastName,
	})

	if err!=nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}

	utils.WriteJSON(w,http.StatusCreated,fmt.Sprintf("user created with email %s",payload.Email))


}
