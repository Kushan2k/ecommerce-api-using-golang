package user

import (
	"net/http"

	"github.com/gorilla/mux"
)


type UserService struct {}



func NewUserService() *UserService {
	return &UserService{}
}


func (s *UserService) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/test",s.Test).Methods("GET")
}

func (s *UserService) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type","application/json")
	w.Write([]byte(`{"message":"Hello World"}`))
}
