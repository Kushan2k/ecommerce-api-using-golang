package main

import (
	"fmt"
	"log"

	"github.com/ecom-api/config"
	"github.com/ecom-api/db"
	"github.com/ecom-api/services/user"
	Mysql "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
);

func main(){
	
	db,err:=db.NewMySqlDatabase(Mysql.Config{
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
	
	server:=fiber.New()
	api:=server.Group("/api/v1/")


	//user router 
	user_router:=api.Group("/auth/")
	user_service:=user.NewUserService(db)
	user_service.RegisterRoutes(user_router)
	// server := api.NewApiServer(":8080", db)


	//view all routes
	server.Get("/routes", func(c *fiber.Ctx) error {
		routes := server.Stack()
		var routeList []string

		for _, routeGroup := range routes {
			for _, route := range routeGroup {
				routeList = append(routeList, fmt.Sprintf("%s %s", route.Method, route.Path))
			}
		}

		return c.JSON(routeList)
	})

	if err:=server.Listen(":8080"); err!=nil{
		log.Fatalf("Error starting server: %s", err)
	}
}
