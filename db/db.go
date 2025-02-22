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

	db,err:=gorm.Open(mysql.Open(cfg.FormatDSN()),&gorm.Config{
		PrepareStmt: true,
		
	})

	if err!=nil{
		return nil,err
	}

	//migration of the models
	db.AutoMigrate(&models.User{},&models.UserAddress{},&models.Category{},&models.Product{},&models.ProductVariation{},&models.Order{},&models.OrderItem{},&models.Cart{},&models.Shipping{},&models.Payment{},&models.Review{},&models.Wishlist{},&models.VariationAttribute{})
	return db, nil
}
