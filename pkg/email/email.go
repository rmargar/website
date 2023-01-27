package email

import (
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

type (
	emailSender func(string, smtp.Auth, string, []string, []byte) error
)

func SendEmail(cfg *SmtpConfig, to []string, message string, send emailSender) error {
	// Send actual message
	err := send(cfg.GetAddress(), cfg.NewAuth(), cfg.Email, to, []byte(message))
	if err != nil {
		log.WithError(err).Error("Email couldn't be sent")
	}
	return err
}
