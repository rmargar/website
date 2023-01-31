package email

import (
	"errors"
	"net/smtp"
	"testing"
)

func TestSendEmailReturnsError(t *testing.T) {
	mockError := errors.New("Oh oh")
	send := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return mockError
	}
	cfg := SmtpConfig{}
	to := []string{"test@host.com"}
	message := []byte("testMsg")
	err := SendEmail(&cfg, to, message, send)

	if err != mockError {
		t.Errorf("expected error (%v), got error (%v)", mockError, err)
	}
}

func TestSendEmailSuccess(t *testing.T) {
	send := func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
		return nil
	}
	cfg := SmtpConfig{}
	to := []string{"test@host.com"}
	message := []byte("testMsg")
	err := SendEmail(&cfg, to, message, send)

	if err != nil {
		t.Errorf("expected error (%v), got error (%v)", nil, err)
	}
}
