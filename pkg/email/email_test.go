package email

import (
	"bytes"
	"errors"
	"net/smtp"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestSendEmail_LogsError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	mockError := errors.New("Oh oh")
	send := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return mockError
	}
	cfg := SmtpConfig{}
	to := []string{"test@host.com"}
	message := []byte("testMsg")
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	SendEmail(&cfg, to, message, send)
	msg := buf.String()
	t.Log(msg)
	if !(strings.Contains(msg, "Email couldn't be sent") && strings.Contains(msg, "level=error")) {
		t.Errorf("SendEmail() did not log error, got: %s", msg)
	}
}

func TestSendEmail_Success(t *testing.T) {
	send := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}
	cfg := SmtpConfig{}
	to := []string{"test@host.com"}
	message := []byte("testMsg")
	SendEmail(&cfg, to, message, send)
}
