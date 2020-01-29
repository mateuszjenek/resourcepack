package services

import (
	"fmt"
	"github.com/jenusek/resourcepack/internal/config"
	"github.com/jenusek/resourcepack/internal/models"
	"net/smtp"
	"strings"
)

func SendMail(to []string, subject string, message string) error {
	auth := smtp.PlainAuth(
		"",
		config.EmailServerAuth.Username,
		config.EmailServerAuth.Password,
		config.EmailServerAuth.SMTPServer,
	)

	msg := "From: " + config.EmailServerAuth.Username + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" + message

	return smtp.SendMail(config.EmailServerAuth.SMTPServer+":587", auth, config.EmailServerAuth.Username, to, []byte(msg))
}

func SendRegisterMail(sender *models.User, dest *models.User, password string) error {
	subject := fmt.Sprintf("Congratulations! %s created account for you on ResourcePack.", sender.Username)
	message := fmt.Sprintf("Credetials to access to service:\nUsername: %s\nPassword: %s", dest.Username, password)

	return SendMail([]string{dest.Email}, subject, message)
}
