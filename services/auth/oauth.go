package auth

import (
	"context"
	"encoding/json"

	"github.com/ecom-api/config"
	"github.com/ecom-api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

var googleOAuthConfig = &oauth2.Config{
	ClientID:     config.Envs.GOOGLE_CLIENT_ID,
	ClientSecret: config.Envs.GOOGLE_CLIENT_SECRET,
	RedirectURL:  "http://localhost:8080/api/v1/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

var oauthStateString = config.Envs.AUTH_SECRET // Prevents CSRF Attacks

type OAuthService struct{
	db *gorm.DB
}

func NewOAuthService(database *gorm.DB) *OAuthService {
	return &OAuthService{
		db: database,
	}
}
// Google Login Route
func (s *OAuthService)  GoogleLogin(c *fiber.Ctx) error {
	url := googleOAuthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// Google Callback Handler
func (s *OAuthService) GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != oauthStateString {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OAuth state"})
	}

	code := c.Query("code")
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// Fetch user info from Google
	client := googleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	// Parse user data
	var userInfo struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse user info"})
	}

	// Check if user exists, else create new user
	var user models.User
	result := s.db.Where("email = ?", userInfo.Email).First(&user)
	if result.Error != nil {
		user = models.User{
			Email:     userInfo.Email,
			FirstName: userInfo.Name,
			Role:      "customer",
		}
		s.db.Create(&user)
	}

	// Generate JWT Token
	tokenString, err := GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT"})
	}

	// Respond with JWT token
	return c.JSON(fiber.Map{"token": tokenString})
}
