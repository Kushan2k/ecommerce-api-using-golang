package main

import (
	"log"

	"github.com/ecom-api/cmd/api"
	"github.com/ecom-api/config"
	"github.com/ecom-api/db"
	"github.com/go-sql-driver/mysql"
);

func main(){
	
	db,err:=db.NewMySqlDatabase(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPass,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})

	if err!=nil{
		log.Fatal("Error connecting to database: ", err)
		return
	}
	
	server := api.NewApiServer(":8080", db)


	if err:=server.Run(); err!=nil{
		log.Fatalf("Error starting server: %s", err)
	}
}
