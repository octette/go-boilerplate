package store

import (
	"context"

	"github.com/ufukomer/go-boilerplate/model"
)

type Store interface {
	GetUserList() ([]*model.User, error)
	GetUser(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(id uint) error
}

// GetUserList gets a list of all users in the system.
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

// GetUser gets the user object corresponding to the given id.
func GetUser(c context.Context, id uint) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

func CreateUser(c context.Context, user *model.User) error {
	return FromContext(c).CreateUser(user)
}

func DeleteUser(c context.Context, id uint) error {
	return FromContext(c).DeleteUser(id)
}
