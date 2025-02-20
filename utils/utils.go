package utils

import (
	"github.com/gofiber/fiber/v2"
)



func WriteJSON(c *fiber.Ctx, status int,v any) error {
	c.Status(status)
	c.Set("Content-Type", "application/json")
	return c.JSON(v)
}

func WriteError(c *fiber.Ctx, status int, err error) error {
	
	return WriteJSON(c,status,map[string]string{"error":err.Error()})
}
