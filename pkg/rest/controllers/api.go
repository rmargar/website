package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/smtp"

	"github.com/rmargar/website/pkg/email"
	e "github.com/rmargar/website/pkg/email"
	log "github.com/sirupsen/logrus"
)

type ContactForm struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Email   string `json:"email"`
}

func HandlePostForm(cfg *email.SmtpConfig, parseForm func(r *http.Request, w http.ResponseWriter) (ContactForm, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contactDetails, err := ParseForm(r, w)
		if err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
		}
		err = e.SendEmail(cfg, []string{"a", "b"}, contactDetails.Message, smtp.SendMail)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := make(map[string]int)
			resp["errorCode"] = http.StatusInternalServerError
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}
		json.NewDecoder(r.Body).Decode(&contactDetails)

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
	if !emailExists || !nameExists || !messageExists {
		return ContactForm{}, errors.New("Error while parsing the form")
	}
	contactDetails := ContactForm{Name: name[0], Email: email[0], Message: message[0]}

	defer r.Body.Close()
	return contactDetails, nil
}
