package representations

import (
	"reflect"
	"testing"
)

type mockInputs struct {
	ID      string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname" validate:"required"`
	Address string `json:"address"`
	Email   string `json:"email" validate:"contains=@"`
}

func TestValidateInputs(t *testing.T) {

	tests := []struct {
		name            string
		args            mockInputs
		wantBool        bool
		wantClientError ClientError
	}{
		{
			name:            "Should return no error",
			args:            mockInputs{ID: "abc", Name: "TestName", Surname: "TestSurname", Address: "TestAddress", Email: "test@test.com"},
			wantBool:        true,
			wantClientError: ClientError{},
		},
		{
			name:     "Should return surname is required",
			args:     mockInputs{ID: "abc", Name: "TestName", Address: "TestAddress", Email: "test@test.com"},
			wantBool: false,
			wantClientError: ClientError{
				Title:  "POST Payload is invalid",
				Code:   "invalid-payload",
				Detail: []string{"surname is required"},
			},
		},
		{
			name:     "Should return surname is required",
			args:     mockInputs{ID: "abc", Name: "TestName", Surname: "TestSurname", Address: "TestAddress", Email: "WrongEmail"},
			wantBool: false,
			wantClientError: ClientError{
				Title:  "POST Payload is invalid",
				Code:   "invalid-payload",
				Detail: []string{"email is invalid"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateInputs(tt.args)
			if got != tt.wantBool {
				t.Errorf("ValidateInputs() got = %v, want %v", got, tt.wantBool)
			}
			if !reflect.DeepEqual(got1, tt.wantClientError) {
				t.Errorf("ValidateInputs() got1 = %v, want %v", got1, tt.wantClientError)
			}
		})
	}
}
