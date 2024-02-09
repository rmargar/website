package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/rmargar/website/pkg/email"
	e "github.com/rmargar/website/pkg/email"
	log "github.com/sirupsen/logrus"
)

type ContactForm struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Email     string `json:"email"`
	WithPhone bool   `json:"withPhone"`
}

func HandlePostForm(cfg *email.SmtpConfig, parseForm func(r *http.Request, w http.ResponseWriter) (ContactForm, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contactDetails, err := ParseForm(r, w)
		if err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
		}
		if contactDetails.WithPhone {
			http.Error(w, "You are a bot", http.StatusBadRequest)
		}
		to := []string{contactDetails.Email, cfg.Email}

		message := []byte(
			fmt.Sprintf("From: %s\r\n", cfg.Email) +
				"Subject: Contact through rmargar.net\r\n\r\n" +
				fmt.Sprintf("Email: %s \n Message: %s\r\n", contactDetails.Email, contactDetails.Message))

		go e.SendEmail(cfg, to, message, smtp.SendMail)

		http.ServeFile(w, r, "./static/submit.html")
	}
}

func ParseForm(r *http.Request, w http.ResponseWriter) (ContactForm, error) {
	err := r.ParseForm()
	if err != nil {
		log.WithError(err).Println("Error while parsing the form")
		return ContactForm{}, errors.New("Error while parsing the form")
	}
	email, emailExists := r.Form["email"]
	name, nameExists := r.Form["name"]
	message, messageExists := r.Form["message"]
	_, withPhone := r.Form["phone"]
	if !emailExists || !nameExists || !messageExists {
		return ContactForm{}, errors.New("Error while parsing the form")
	}
	contactDetails := ContactForm{Name: name[0], Email: email[0], Message: message[0], WithPhone: withPhone}

	defer r.Body.Close()
	return contactDetails, nil
}
