package model

import (
	"testing"

	"github.com/elgs/gostrgen"
)

func TestUserValidate(t *testing.T) {
	str, err := gostrgen.RandGen(256, gostrgen.Lower, "abcdefghijklmnoprstuvyz", "")
	if err != nil {
		t.Errorf("Long string generation error: %s", err.Error())
	}

	var tests = []struct {
		user User
		err  error
	}{
		{
			user: User{},
			err:  errUserEmailInvalid,
		},
		{
			user: User{Email: ""},
			err:  errUserEmailInvalid,
		},
		{
			user: User{Email: "ufukomer"},
			err:  errUserEmailInvalid,
		},
		{
			user: User{Email: "!ufukomer@gmail.com"},
			err:  errUserEmailInvalid,
		},
		{
			user: User{Email: str + "ufukomer@gmail.com"},
			err:  errUserEmailInvalid,
		},
		{
			user: User{Email: "ufukomer@gmail.com"},
			err:  nil,
		},
		{
			user: User{Email: "onerciller@gmail.com"},
			err:  nil,
		},
	}

	for _, test := range tests {
		err := test.user.Validate()
		if want, got := test.err, err; want != got {
			t.Errorf("Want user validation error %v, got %v", want, got)
		}
	}
}
