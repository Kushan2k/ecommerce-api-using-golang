package utils

import (
	"crypto/tls"

	"github.com/ecom-api/config"
	"gopkg.in/gomail.v2"
)


func GetMailer() *gomail.Dialer {

	d:=gomail.NewDialer(config.Envs.MailHost,config.Envs.MailPort,config.Envs.MailUser,config.Envs.MailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d
}

func GetMessage() *gomail.Message {
	return gomail.NewMessage()
}
