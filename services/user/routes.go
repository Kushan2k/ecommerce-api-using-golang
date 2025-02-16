package user

import (
	"net/http"

	"github.com/ecom-api/types"
	"github.com/ecom-api/utils"
	"github.com/gorilla/mux"
)


type UserService struct {}



func NewUserService() *UserService {
	return &UserService{}
}


func (s *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/test",s.Test).Methods("GET")
	router.HandleFunc("/register",s.RegisterUser).Methods("POST")
}

func (s *UserService) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type","application/json")
	w.Write([]byte(`{"message":"Hello World"}`))
}


func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	var payload types.RegisterBodyType

	err:=utils.ParseJSON(r,payload)

	if err!=nil{
		utils.WriteError(w,http.StatusBadRequest,err)
	}

	utils.WriteJSON(w,http.StatusOK,payload)


}
