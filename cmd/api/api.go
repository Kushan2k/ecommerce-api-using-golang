package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ecom-api/services/user"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
	db *sql.DB
}



func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db: db,
	}
}

func (s *ApiServer) Run() error {

	router:=mux.NewRouter();
	subRouter:=router.PathPrefix("/api/v1").Subrouter();
	_ = subRouter;

	usestore:=user.NewStore(s.db)

	user_service:=user.NewUserService(usestore)
	user_service.RegisterRoutes(subRouter)
	
	log.Println("Server is running on port 8080")
	return http.ListenAndServe(s.addr, nil);
}
