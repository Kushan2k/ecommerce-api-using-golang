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
	router.Get("/",s.get_all_categories)
	router.Get("/:id",s.get_category_by_id)
	
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

func (s *CategoryService) get_all_categories(c *fiber.Ctx) error{
	var categories []models.Category
	if err:=s.db.Find(&categories).Error;err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":err.Error(),
		})
	}
	return c.Status(200).JSON(categories)
}


func (s *CategoryService) get_category_by_id(c *fiber.Ctx) error{
	id:=c.Params("id","")

	if id==""{
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid category id",
		})
	}

	var category models.Category

	res:=s.db.Model(&models.Category{}).Where("id = ?",id).First(&category)

	if res.Error!=nil{
		return c.Status(200).JSON(category)
	}

	return c.Status(404).JSON(fiber.Map{
			"error": "Category not found",
		})

	
}
