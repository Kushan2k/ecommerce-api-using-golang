package auth

import (
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

func GenerateJWT(vars jwt.MapClaims) (string,error){
	key:=config.Envs.JWT_KEY

	t:=jwt.NewWithClaims(jwt.SigningMethodHS256,vars)

	s,err:=t.SignedString([]byte(key))

	if err!=nil{
		return "",err
	}

	return s,nil


}
