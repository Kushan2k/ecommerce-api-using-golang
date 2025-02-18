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

	MailHost string
	MailUser string
	MailPass string
	MailPort int
	


}

var Envs=initConfig()

func initConfig() Config {

	godotenv.Load()

	// $mail->Host       = 'smtp.titan.email'; //Set the SMTP server to send through
  //   $mail->SMTPAuth   = true;    //Enable SMTP authentication
  //   $mail->Username   = 'testadmin@forgear.edu.lk';  //SMTP username
  //   $mail->Password   = "7>6{SW#N('> &Do";     //SMTP password
  //   $mail->SMTPSecure = PHPMailer::ENCRYPTION_SMTPS;   //Enable implicit TLS encryption
  //   $mail->Port       = 465; 

  //   //Recipients
  //   $mail->setFrom('testadmin@forgear.edu.lk', 'Bill Remailder');
	return Config{
		PublicHost: getENV("PUBLIC_HOST", "localhost"),
		Port: getENV("PORT", "8080"),
		DBUser: getENV("DB_USER", "root"),
		DBPass: getENV("DB_PASS", ""),
		DBAddress: getENV("DB_ADDRESS", "localhost:3306"),
		DBName: getENV("DB_NAME", "go_ecom_api"),

		MailHost: getENV("MAIL_HOST", "smtp.titan.email"),
		MailUser: getENV("MAIL_USER", ""),
		MailPass: getENV("MAIL_PASS", ""),
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
	