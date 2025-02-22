package config

import (
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublicHost string
	Port			 string
	DBUser		 string
	DBPass		 string
	DBAddress	string
	DBName		 string
	JWT_KEY string

	MailHost string
	MailUser string
	MailPass string
	MailPort int

	EXPIRE_TIME_MULTIPLER int
	


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
		JWT_KEY: getENV("JWT_SECRET","123"),

		MailHost: getENV("MAIL_HOST", "smtp.titan.email"),
		MailUser: getENV("MAIL_USER", ""),
		MailPass: getENV("MAIL_PASS", ""),
		EXPIRE_TIME_MULTIPLER: func() int {
			multiplier, err := strconv.Atoi(getENV("EXPIRE_TIME_MULTIPLER", "24"))
			if err != nil {
				return 24
			}
			return multiplier
		}(),
		MailPort: func() int {
			port, err := strconv.Atoi(getENV("MAIL_PORT", "465"))
			if err != nil {
				return 465
			}
			return port
		}(),

		

	}
}

func getENV(key,fallback string) string {
	if value,ok:=os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
	