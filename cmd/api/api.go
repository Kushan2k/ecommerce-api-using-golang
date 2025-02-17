package api

import (
	"log"
	"net/http"

	"github.com/ecom-api/services/user"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ApiServer struct {
	addr string
	db *gorm.DB
}



func NewApiServer(addr string, db *gorm.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db: db,
	}
}

func (s *ApiServer) Run() error {

	router:=mux.NewRouter();
	subRouter:=router.PathPrefix("/api/v1").Subrouter();
	// usestore:=user.NewStore(s.db)
	subRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	user_service:=user.NewUserService(s.db)
	user_service.RegisterRoutes(subRouter)
	
	log.Printf("Server is running on port %s",s.addr)
	return http.ListenAndServe(s.addr, nil);
}
