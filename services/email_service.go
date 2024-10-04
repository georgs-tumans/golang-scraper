package services

import (
	"fmt"
	"scraper/config"

	"github.com/go-mail/mail"
)

type EmailClient struct {
	SMTPHost   string
	Port       int
	Password   string
	Username   string
	Recipients []string
}

func NewEmailClient() *EmailClient {
	config := config.GetConfig()

	return &EmailClient{
		SMTPHost:   config.SMPTHost,
		Port:       config.SMTPPort,
		Password:   config.SMTPPassword,
		Username:   config.SMTPUsername,
		Recipients: config.EmailRecipients,
	}
}

func (c *EmailClient) SendEmail(subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", c.Username)
	m.SetHeader("To", c.Recipients...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := mail.NewDialer(c.SMTPHost, c.Port, c.Username, c.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("[SendEmail] Error sending email", err)

		return err
	}

	return nil
}
