package controllers

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFormSucces(t *testing.T) {
	assert := assert.New(t)

	reader := strings.NewReader("email=hello@mail.com&message=test&name=testName")
	request := httptest.NewRequest("POST", "http://localhost/", reader)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	writer := httptest.NewRecorder()

	contactDetails, err := ParseForm(request, writer)

	expectedContactDetails := ContactForm{"testName", "test", "hello@mail.com"}
	assert.Equal(expectedContactDetails, contactDetails)

	if err != nil {
		t.Errorf("expected error (%v), got error (%v)", nil, err)
	}
}

func TestParseFormError(t *testing.T) {
	assert := assert.New(t)

	writer := httptest.NewRecorder()

	subtests := []struct {
		name   string
		header string
		reader *strings.Reader
	}{
		{
			name:   "wrong header",
			header: "application/json",
			reader: strings.NewReader("email=hello@mail.com&message=test&name=testName"),
		},
		{
			name:   "incomplete form",
			header: "application/x-www-form-urlencoded",
			reader: strings.NewReader("message=test&name=testName"),
		},
	}
	for _, subtest := range subtests {
		request := httptest.NewRequest("POST", "http://localhost/", subtest.reader)
		request.Header.Set("Content-Type", subtest.header)
		test := func(t *testing.T) {
			contactDetails, err := ParseForm(request, writer)
			assert.Equal(ContactForm{}, contactDetails)
			expectedError := errors.New("Error while parsing the form")
			if err.Error() != expectedError.Error() {
				t.Errorf("expected error (%v), got error (%v)", expectedError, err)
			}
		}
		t.Run(subtest.name, test)
	}
}
