package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/ecom-api/config"
	"github.com/ecom-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/golang-jwt/jwt/v5"
)

func Is_authenticated(c *fiber.Ctx) error {
    // Check if user is authenticated
    authHeader := c.Get("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("unauthorized"))
    }

    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrSignatureInvalid
        }
        return []byte(config.Envs.JWT_KEY), nil
    })

    if err != nil {
        return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("token parsing error: %v", err))
    }

    if !token.Valid {
        return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("token is invalid"))
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("failed to decode the claims"))
    }

    fmt.Println("Decoded Claims:", claims) // Debug print

		// Debug print

    userID,found:= claims["user_id"]
    if !found {
			fmt.Println("User ID:", userID)
        return utils.WriteError(c, fiber.StatusUnauthorized, fmt.Errorf("user_id not found in claims"))
    }

    fmt.Println("User ID:", userID) // Debug print
    c.Locals("user_id", userID)

    return c.Next()
}


func Limiter(c *fiber.Ctx) error {
    return limiter.New(
        limiter.Config{
            Max: 6,
            Expiration: 5 * time.Minute,
            SkipFailedRequests: true,
        },
    )(c)
}
