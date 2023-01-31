package email

import (
	"fmt"
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

type (
	emailSender func(string, smtp.Auth, string, []string, []byte) error
)

func SendEmail(cfg *SmtpConfig, to []string, message []byte, send emailSender) error {
	log.Info("Sending email to", to)
	// Send actual message
	err := send(cfg.GetAddress(), cfg.NewAuth(), cfg.Email, to, message)
	if err != nil {
		fmt.Println(to)
		log.WithError(err).Error("Email couldn't be sent")
		return err
	}
	log.Info("Email sent succesfully")
	return nil
}
