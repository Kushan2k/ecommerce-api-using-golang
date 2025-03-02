package shop

import (
	"github.com/ecom-api/middlewares"
	"github.com/ecom-api/models"
	"github.com/ecom-api/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type ShopService struct {
	db *gorm.DB
}


func NewShopService(db *gorm.DB) *ShopService {
	return &ShopService{db}
}


func (s *ShopService) RegisterRoutes(router fiber.Router)  {
	//public routes
	router.Get("/",s.get_all_shops)
	router.Get("/:id",s.get_shop_by_id)

	//authenticated routes
	router.Post("/",middlewares.Is_authenticated,s.CreateShop)
	router.Put("/:id",middlewares.Is_authenticated,s.UpdateShop)
}

func (s *ShopService) CreateShop(c *fiber.Ctx) error {
	
	var shop types.ShopRequest

	if err:=c.BodyParser(&shop);err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	tx:=s.db.Create(&shop)

	if tx.Error!=nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Error creating shop",
		})
	}

	return c.Status(201).JSON(shop)
}

func (s *ShopService) UpdateShop(c *fiber.Ctx) error {

	id:=c.Params("id","")

	if id==""{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid shop id",
		})
	}

	var shop types.ShopRequest

	if err:=c.BodyParser(&shop);err!=nil{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	var shop_model models.Shop

	s.db.Model(&models.Shop{}).Where("id = ?",id).First(&shop_model)

	shop_model.ShopName=shop.Name
	shop_model.ShopDescription=shop.Description
	shop_model.ShopLogoURL=&shop.LogoURL

	tx:=s.db.Save(&shop_model)

	if tx.Error!=nil{
		return c.Status(500).JSON(fiber.Map{
			"error": "Error updating shop",
		})
	}

	return c.Status(200).JSON(shop_model)

}


func (s *ShopService) get_shop_by_id(c *fiber.Ctx) error {
	
	id:=c.Params("id","")

	if id==""{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid shop id",
		})
	}

	var shop models.Shop

	s.db.Model(&models.Shop{}).Where("id = ?",id).Preload("Products").First(&shop)

	return c.JSON(shop)
}

func (s *ShopService) get_all_shops(c *fiber.Ctx) error {
	
	var shops []models.Shop

	s.db.Model(&models.Shop{}).Find(&shops)

	return c.JSON(shops)
}
