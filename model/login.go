package model

import (
	"errors"
)

// email will be validated under user model

// validate password
var errLoginPasswordInvalid = errors.New("Invalid Password")

// Login structure.
type Login struct {
	Email    string
	Password string
}

// Validate validates the required fields and formats.
func (l *Login) Validate() error {
	switch {
	case len(l.Password) < 8:
		return errLoginPasswordInvalid
	case len(l.Password) > 250:
		return errLoginPasswordInvalid
	default:
		return nil
	}
}
