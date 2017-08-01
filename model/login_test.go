package model

import (
	"testing"

	"github.com/elgs/gostrgen"
)

func TestLoginValidate(t *testing.T) {
	str, err := gostrgen.RandGen(256, gostrgen.Digit, "0123456789", "")
	if err != nil {
		t.Errorf("Long string generation error: %s", err.Error())
	}

	var tests = []struct {
		login Login
		err   error
	}{
		{
			login: Login{},
			err:   errLoginPasswordInvalid,
		},
		{
			login: Login{Password: "1234567"},
			err:   errLoginPasswordInvalid,
		},
		{
			login: Login{Password: str},
			err:   errLoginPasswordInvalid,
		},
		{
			login: Login{Password: "12345678"},
			err:   nil,
		},
		{
			login: Login{Password: "123456!%"},
			err:   nil,
		},
	}

	for _, test := range tests {
		err := test.login.Validate()
		if want, got := test.err, err; want != got {
			t.Errorf("Want login validation error %v, got %v", want, got)
		}
	}
}
