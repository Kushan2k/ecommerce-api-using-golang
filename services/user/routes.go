package user

import (
	"net/http"

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

	

	utils.WriteJSON(w,http.StatusOK,payload)


}
