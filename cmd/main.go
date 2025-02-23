package main

import (
	"fmt"
	"log"

	"github.com/ecom-api/config"
	"github.com/ecom-api/db"
	"github.com/ecom-api/services/product"
	"github.com/ecom-api/services/user"
	Mysql "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	//global middlewares
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	server.Use(helmet.New())
	server.Use(logger.New())


	//user router 
	user_router:=api.Group("/auth/")
	user_service:=user.NewUserService(db)
	user_service.RegisterRoutes(user_router)
	// server := api.NewApiServer(":8080", db)

	//product routes
	product_router:=api.Group("/products/")
	product_service:=product.NewProductService(db)
	product_service.RegisterRoutes(product_router)


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
