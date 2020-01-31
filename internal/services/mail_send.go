package services

import (
	"fmt"
	"github.com/jenusek/resourcepack/internal/models"
	"net/smtp"
	"strings"
)

func sendMail(config *models.Configuration, to []string, subject string, message string) error {
	auth := smtp.PlainAuth(
		"",
		config.EmailServer.Username,
		config.EmailServer.Password,
		config.EmailServer.SMTP,
	)

	msg := "From: " + config.EmailServer.Username + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + subject + "\n\n" + message

	return smtp.SendMail(config.EmailServer.SMTP+":587", auth, config.EmailServer.Username, to, []byte(msg))
}

// SendRegisterMail sends mail to dest User with credentials to log in to server
func SendRegisterMail(config *models.Configuration, sender *models.User, dest *models.User, password string) error {
	subject := fmt.Sprintf("Congratulations! %s created account for you on ResourcePack.", sender.Username)
	message := fmt.Sprintf("Credetials to access to service:\nUsername: %s\nPassword: %s", dest.Username, password)

	return sendMail(config, []string{dest.Email}, subject, message)
}
