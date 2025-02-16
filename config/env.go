package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port			 string
	DBUser		 string
	DBPass		 string
	DBAddress	string
	DBName		 string

}

var Envs=initConfig()

func initConfig() Config {

	godotenv.Load()

	return Config{
		PublicHost: getENV("PUBLIC_HOST", "localhost"),
		Port: getENV("PORT", "8080"),
		DBUser: getENV("DB_USER", "root"),
		DBPass: getENV("DB_PASS", ""),
		DBAddress: getENV("DB_ADDRESS", "localhost:3306"),
		DBName: getENV("DB_NAME", "go_ecom_api"),
	}
}

func getENV(key,fallback string) string {
	if value,ok:=os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
	