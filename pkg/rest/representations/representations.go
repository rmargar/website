package representations

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type ClientError struct {
	Title  string   `json:"title"`
	Detail []string `json:"detail"`
	Code   string   `json:"code"`
	Source string   `json:"source"`
}

type RepresentationForClientErrors struct {
	Errors []ClientError `json:"errors,omitempty"`
}

var ErrInvalidPostPayloadJSON = func(detail []string) ClientError {
	return ClientError{
		Title:  "POST Payload is invalid",
		Code:   "invalid-payload",
		Detail: detail,
	}
}

func WriteBadRequestWithErr(w http.ResponseWriter, err error) {
	log.WithError(err).Warn("Bad request error")
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func ValidateInputs(data interface{}) (bool, ClientError) {
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			log.WithError(err).Warn("Invalid validation error")
		}

		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(data)

		detailedErrors := []string{}
		for _, err := range err.(validator.ValidationErrors) {

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				detailedErrors = append(detailedErrors, name+" is required")
				break
			default:
				detailedErrors = append(detailedErrors, name+" is invalid")
				break
			}
		}
		return false, ErrInvalidPostPayloadJSON(detailedErrors)
	}
	return true, ClientError{}
}

func WriteValidationResponse(errors []ClientError, writer http.ResponseWriter) {

	var errorResponse RepresentationForClientErrors

	for _, err := range errors {
		errorResponse.Errors = append(errorResponse.Errors, err)
	}
	message, err := json.Marshal(errorResponse)

	if err != nil {
		//An error occurred processing the json
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("An error occured internally"))
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	writer.Write(message)
}
