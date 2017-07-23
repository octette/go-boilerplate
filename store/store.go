package store

import (
	"context"

	"github.com/ufukomer/go-boilerplate/model"
)

type Store interface {
	GetUserList() ([]*model.User, error)
	GetUser(id int64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

// GetUserList gets a list of all users in the system.
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

func GetUser(c context.Context, id int64) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

func GetUserByEmail(c context.Context, email string) (*model.User, error) {
	return FromContext(c).GetUserByEmail(email)
}
