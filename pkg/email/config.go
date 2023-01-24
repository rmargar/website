package email

import (
	"net/smtp"
)

type SmtpConfig struct {
	Host     string `env:"SFTP_HOST"`
	Username string `env:"SFTP_USERNAME"`
	Password string `env:"SFTP_PASSWORD"`
	Port     string `env:"SFTP_PORT"`
	Email    string `env:"SFTP_EMAIL"`
}

func (cfg *SmtpConfig) GetAddress() string {
	return cfg.Host + ":" + cfg.Port
}

func (cfg *SmtpConfig) NewAuth() smtp.Auth {
	return smtp.PlainAuth(cfg.Email, cfg.Username, cfg.Password, cfg.Host)
}
