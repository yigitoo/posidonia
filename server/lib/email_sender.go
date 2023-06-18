// THIS CODE IS DEPRECATED
package lib

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/BinodKafle/gomail/gomail"
)

func EmailSender(type_of_mail_protocol string) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	params := os.Args
	paramsLength := len(params)
	if paramsLength < 2 {
		log.Println("Please add SMTP or OAUTH along with go run main.go command")
		log.Println("Eg: go run main.go SMTP or go run main.go OAUTH")
		os.Exit(1)
	}

	inputMethod := type_of_mail_protocol

	valid := IsValidInputMethod(inputMethod)

	emailTo := "gumusyigit101@gmail.com"

	if valid {
		data := struct {
			ReceiverName string
			SenderName   string
		}{
			ReceiverName: "Yiğit GÜMÜŞ",
			SenderName:   "Yiğit GÜMÜŞ",
		}

		if inputMethod == "SMTP" {
			status, err := gomail.SendEmailSMTP([]string{emailTo}, data, "sample_template.txt")
			if err != nil {
				log.Println(err)
			}
			if status {
				log.Println("Email sent successfully using SMTP")
			}
		}

		if inputMethod == "OAUTH" {
			gomail.OAuthGmailService()
			status, err := gomail.SendEmailOAUTH2(emailTo, data, "sample_template.txt")
			if err != nil {
				log.Println(err)
			}
			if status {
				log.Println("Email sent successfully using OAUTH")
				return nil
			}
		}
	} else {
		log.Println("Please add SMTP or OAUTH along with go run main.go command")
		log.Println("Eg: go run main.go SMTP or go run main.go OAUTH")
		os.Exit(1)
		return errors.New("Error: email not sended")
	}
	return nil
}

func IsValidInputMethod(method string) bool {
	switch method {
	case
		"SMTP",
		"OAUTH":
		return true
	}
	return false
}
