package product

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



type ProductService struct {
	db *gorm.DB
}


func NewProductService(database *gorm.DB) *ProductService {
	return &ProductService{
		db: database,
	}
}


func (s *ProductService) RegisterRoutes(router fiber.Router) {
	router.Get("/",s.get_all_products)
	router.Get("/:id",s.get_product_by_id)
	router.Post("/",s.create_product)
	router.Put("/:id",s.update_product)
	router.Delete("/:id",s.delete_product)
}

func (* ProductService) get_all_products(c *fiber.Ctx) error{
	return c.JSON("All products")
}

func (* ProductService) get_product_by_id(c *fiber.Ctx) error{
	return c.JSON("Product by id")
}

func (* ProductService) create_product(c *fiber.Ctx) error{
	return c.JSON("Create product")
}

func (* ProductService) update_product(c *fiber.Ctx) error{
	return c.JSON("Update product")
}

func (* ProductService) delete_product(c *fiber.Ctx) error{
	return c.JSON("Delete product")
}
