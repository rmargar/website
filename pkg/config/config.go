package config

import (
	log "github.com/sirupsen/logrus"

	cleanenv "github.com/ilyakaznacheev/cleanenv"
	"github.com/rmargar/website/pkg/database"
	"github.com/rmargar/website/pkg/email"
)

type Config struct {
	Port       string `env:"PORT" env-default:"8000"`
	JwtSecret  string `env:"JWT_SECRET"`
	SmtpConfig email.SmtpConfig
	Database   database.DatabaseConfig
}

func GetConfig() *Config {

	var cfg Config
	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		log.WithError(err).Fatal("Couldn't read the environment")
	}
	return &cfg
}
