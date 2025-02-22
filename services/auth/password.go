package auth

import (
	"time"

	"github.com/ecom-api/config"
	"github.com/ecom-api/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string,error){
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err!=nil{
		return "",err
	}

	return string(hash),nil
}

func CheckPasswordHash(password,hash string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))

	return err==nil
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(user models.User) (string, error) {
    key := config.Envs.JWT_KEY

		claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Envs.EXPIRE_TIME_MULTIPLER) * time.Hour)), // Token expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ecommerce-api",
		},
	}

    // Ensure the correct HMAC algorithm is used
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the token with the HMAC key
    s, err := t.SignedString([]byte(key))
    if err != nil {
        return "", err
    }

    return s, nil
}
