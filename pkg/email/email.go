package email

import (
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

func SendEmail(cfg *SmtpConfig, to []string, message string) error {
	// Send actual message
	err := smtp.SendMail(cfg.GetAddress(), cfg.NewAuth(), cfg.Email, to, []byte(message))
	if err != nil {
		log.WithError(err).Error("Email couldn't be sent")
	}
	return err
}
