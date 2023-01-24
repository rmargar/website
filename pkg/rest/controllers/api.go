package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rmargar/website/pkg/email"
	e "github.com/rmargar/website/pkg/email"
)

type ContactForm struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Email   string `json:"email"`
}

func HandlePostForm(cfg *email.SmtpConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.ParseForm() != nil {
			fmt.Println("error while parsing the form")
		}
		email, emailExists := r.Form["email"]
		name, nameExists := r.Form["name"]
		message, messageExists := r.Form["message"]
		if !emailExists || !nameExists || !messageExists {
			http.Error(w, "Invalid form", http.StatusBadRequest)
		}
		contactDetails := ContactForm{Name: name[0], Email: email[0], Message: message[0]}

		defer r.Body.Close()
		err := e.SendEmail(cfg, []string{"a", "b"}, message[0])
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
