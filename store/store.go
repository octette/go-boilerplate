package store

import (
	"context"

	"github.com/ufukomer/tagon-api/model"
)

type Store interface {
	GetUserList() ([]*model.User, error)
}

// GetUserList gets a list of all users in the system.
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}
