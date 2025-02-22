package db

import (
	"fmt"

	"github.com/ecom-api/models"
	Mysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlDatabase(cfg Mysql.Config) (*gorm.DB, error) {

	fmt.Println(cfg.FormatDSN())

	db,err:=gorm.Open(mysql.Open(cfg.FormatDSN()))

	if err!=nil{
		return nil,err
	}

	//migration of the models
	db.AutoMigrate(&models.User{},&models.Shop{},&models.Product{})
	return db, nil
}
