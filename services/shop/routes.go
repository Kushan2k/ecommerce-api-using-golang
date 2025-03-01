package shop

import (
	"github.com/ecom-api/models"
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
	router.Get("/",s.get_all_shops)
}

func (s *ShopService) get_all_shops(c *fiber.Ctx) error {
	
	var shops []models.Shop

	s.db.Model(&models.Shop{}).Find(&shops)

	return c.JSON(shops)
}
