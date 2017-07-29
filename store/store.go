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
	GetPostList() ([]*model.Post, error)
	GetPost(id uint) (*model.Post, error)
	CreatePost(user *model.Post) error
	UpdatePost(post *model.Post, id uint) (*model.Post, error)
	DeletePost(id uint) error
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

// GetPostList gets a list of all users in the system.
func GetPostList(c context.Context) ([]*model.Post, error) {
	return FromContext(c).GetPostList()
}

// GetUser gets the user object corresponding to the given id.
func GetPost(c context.Context, id uint) (*model.Post, error) {
	return FromContext(c).GetPost(id)
}

func CreatePost(c context.Context, post *model.Post) error {
	return FromContext(c).CreatePost(post)
}

func UpdatePost(c context.Context, post *model.Post, id uint) (*model.Post, error) {
	return FromContext(c).UpdatePost(post, id)
}

func DeletePost(c context.Context, id uint) error {
	return FromContext(c).DeletePost(id)
}
