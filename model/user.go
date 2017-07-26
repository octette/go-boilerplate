package model

import (
	"errors"
	"regexp"

	"github.com/jinzhu/gorm"
)

// validate a email
var reEmail = regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$")

var errUserEmailInvalid = errors.New("Invalid User Email")

type User struct {
	gorm.Model
	Email    string
	Password string
}

// Validate validates the required fields and formats.
func (u *User) Validate() error {
	switch {
	case len(u.Email) == 0:
		return errUserEmailInvalid
	case len(u.Email) > 250:
		return errUserEmailInvalid
	case !reEmail.MatchString(u.Email):
		return errUserEmailInvalid
	default:
		return nil
	}
}
