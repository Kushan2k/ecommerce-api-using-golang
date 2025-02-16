package main

import (
	"log"

	"github.com/ecom-api/cmd/api"
);

func main(){
	server := api.NewApiServer(":8080", nil)



	if err:=server.Run(); err!=nil{
		log.Fatalf("Error starting server: %s", err)
	}
}
