package categories

import (
	"github.com/ecom-api/models"
	"github.com/ecom-api/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(database *gorm.DB) *CategoryService {
	return &CategoryService{
		db: database,
	}
}



func (s *CategoryService) RegisterRoutes(router fiber.Router) {
	router.Post("/",s.create_category)
	
}

func (s *CategoryService) create_category(c *fiber.Ctx) error{
	var category types.CategoryCreateRequest

	if err:=c.BodyParser(&category);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":err.Error(),
		})
	}

	res:=s.db.Where("name=?",category.Name).First(&models.Category{})
	if res.RowsAffected>0{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Category already exists",
		})
	}


	if err:=s.db.Create(&models.Category{
		Name:category.Name,
	}).Error;err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return c.Status(201).JSON(category)
}
