package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/rmargar/website/pkg/config"
	e "github.com/rmargar/website/pkg/email"
	log "github.com/sirupsen/logrus"
)

const siteVerifyURL = "https://www.google.com/recaptcha/api/siteverify"

type ContactForm struct {
	Name              string `json:"name"`
	Message           string `json:"message"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

func HandlePostForm(cfg *config.Config, parseForm func(r *http.Request, w http.ResponseWriter) (ContactForm, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contactDetails, err := ParseForm(r, w)
		if err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
		}
		if contactDetails.Phone != "" || checkRecaptcha(cfg.RecaptchaSecret, contactDetails.RecaptchaResponse) != nil {
			http.Error(w, "You are a bot", http.StatusBadRequest)
		}
		to := []string{contactDetails.Email, cfg.SmtpConfig.Email}

		message := []byte(
			fmt.Sprintf("From: %s\r\n", cfg.SmtpConfig.Email) +
				"Subject: Contact through rmargar.net\r\n\r\n" +
				fmt.Sprintf("Email: %s \n Message: %s\r\n", contactDetails.Email, contactDetails.Message))

		go e.SendEmail(&cfg.SmtpConfig, to, message, smtp.SendMail)

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
	phone, phoneExists := r.Form["phone"]
	recaptchaResponse, recaptchaResponseExists := r.Form["g-recaptcha-response"]
	if !emailExists || !nameExists || !messageExists || !phoneExists || !recaptchaResponseExists {
		return ContactForm{}, errors.New("Error while parsing the form")
	}
	contactDetails := ContactForm{Name: name[0], Email: email[0], Message: message[0], Phone: phone[0], RecaptchaResponse: recaptchaResponse[0]}

	defer r.Body.Close()
	return contactDetails, nil
}

func checkRecaptcha(secret, response string) error {
	req, err := http.NewRequest(http.MethodPost, siteVerifyURL, nil)
	if err != nil {
		return err
	}

	// Add necessary request parameters.
	q := req.URL.Query()
	q.Add("secret", secret)
	q.Add("response", response)
	req.URL.RawQuery = q.Encode()

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response.
	var body SiteVerifyResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return err
	}

	// Check recaptcha verification success.
	if !body.Success {
		return errors.New("unsuccessful recaptcha verify request")
	}

	// Check response action.
	if body.Action != "CONTACT" {
		return errors.New("mismatched recaptcha action")
	}

	return nil
}
