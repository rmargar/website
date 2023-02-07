package email

import (
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

type (
	emailSender func(string, smtp.Auth, string, []string, []byte) error
)

func SendEmail(cfg *SmtpConfig, to []string, message []byte, send emailSender) {
	log.Info("Sending email to", to)
	// Send actual message
	err := send(cfg.GetAddress(), cfg.NewAuth(), cfg.Email, to, message)
	if err != nil {
		log.WithError(err).Error("Email couldn't be sent")
		return
	}
	log.Info("Email sent succesfully")

}
