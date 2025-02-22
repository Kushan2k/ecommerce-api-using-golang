package auth

import (
	"time"

	"github.com/ecom-api/config"
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

func GenerateJWT(vars jwt.MapClaims) (string, error) {
    key := config.Envs.JWT_KEY

		expirationTime := time.Now().Add(time.Duration(config.Envs.EXPIRE_TIME_MULTIPLER) * time.Hour).Unix()

		vars["exp"] = expirationTime
    // Ensure the correct HMAC algorithm is used
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, vars)

    // Sign the token with the HMAC key
    s, err := t.SignedString([]byte(key))
    if err != nil {
        return "", err
    }

    return s, nil
}
