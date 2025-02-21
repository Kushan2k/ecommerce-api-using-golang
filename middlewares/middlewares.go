package middlewares

import (
	"fmt"
	"strings"

	"github.com/ecom-api/config"
	"github.com/ecom-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func  Is_authenticated(c *fiber.Ctx) error{
	//check if user is authenticated

	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("unauthorized"))
		// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return config.Envs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("token invalid"))
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userID, ok := claims["user_id"].(string)

	if !ok {
		return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("invalid token claims"))
	}

	c.Locals("user_id", userID)

	return c.Next()

	

}
