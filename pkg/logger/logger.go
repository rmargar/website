package logging

import (
	log "github.com/sirupsen/logrus"
)

func LogFormatter() *log.JSONFormatter {
	return &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyMsg:  "message",
		},
	}
}

func NewLogger() {
	log.SetFormatter(LogFormatter())
	log.SetLevel(log.DebugLevel)
}
