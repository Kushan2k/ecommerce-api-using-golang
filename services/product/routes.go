package product

import (
	"time"

	"github.com/ecom-api/middlewares"
	"github.com/ecom-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gorm.io/gorm"
)



type ProductService struct {
	db *gorm.DB
}

type QueryParams struct {
	Page int `json:"page"`
	PerPage int `json:"per_page"`
	Search string `json:"search"`
}


func NewProductService(database *gorm.DB) *ProductService {
	return &ProductService{
		db: database,
	}
}


func (s *ProductService) RegisterRoutes(router fiber.Router) {
	router.Get("/",limiter.New(limiter.Config{
            Max: 6,
            Expiration: 10 * time.Second,
        }),s.get_all_products)
				
	router.Get("/:id",s.get_product_by_id)
	router.Post("/",middlewares.Is_authenticated,s.create_product)
	router.Put("/:id",middlewares.Is_authenticated,s.update_product)
	router.Delete("/:id",middlewares.Is_authenticated,s.delete_product)
}

func (s *ProductService) get_all_products(c *fiber.Ctx) error{

	var query QueryParams=QueryParams{
		Page:c.QueryInt("page",1),
		PerPage:c.QueryInt("per_page",10),
		Search:c.Query("search"),
	}


	var products []models.Product
	if query.Search!=""{
		s.db.Find(&products).Where("name LIKE ?", "%"+query.Search+"%").Limit(query.PerPage).Offset((query.Page-1)*query.PerPage).Preload("ProductImages").Preload("Category").Preload("ProductVariation").Preload("VariationAttribute").Preload("VariantImage")

	}else {
		s.db.Find(&products).Where("id > ?",0).Limit(query.PerPage).Offset((query.Page-1)*query.PerPage).Preload("ProductImages").Preload("Category").Preload("ProductVariation").Preload("VariationAttribute").Preload("VariantImage")
	}

	return c.JSON(products)
}

func (s *ProductService) get_product_by_id(c *fiber.Ctx) error{
	return c.JSON("Product by id")
}

func (s *ProductService) create_product(c *fiber.Ctx) error{
	return c.JSON("Create product")
}

func (s *ProductService) update_product(c *fiber.Ctx) error{
	return c.JSON("Update product")
}

func (s *ProductService) delete_product(c *fiber.Ctx) error{
	return c.JSON("Delete product")
}
