package store

import (
	"context"

	"github.com/ufukomer/go-boilerplate/model"
)

type Store interface {
	GetUserList() ([]*model.User, error)
	GetUser(id uint) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByLogin(email string, password string) (*model.User, error)
	CreateUser(user *model.User) error
	DeleteUser(id uint) error
	GetPostList() ([]*model.Post, error)
	GetPost(id uint) (*model.Post, error)
	CreatePost(user *model.Post) error
	UpdatePost(post *model.Post, id uint) (*model.Post, error)
	DeletePost(id uint) error
	GetComment(id uint) (*model.Comment, error)
	CreateComment(comment *model.Comment) error
	UpdateComment(comment *model.Comment, id uint) (*model.Comment, error)
	DeleteComment(id uint) error
}

// GetUserList gets a list of all users in the system.
func GetUserList(c context.Context) ([]*model.User, error) {
	return FromContext(c).GetUserList()
}

// GetUser gets the user object corresponding to the given id.
func GetUser(c context.Context, id uint) (*model.User, error) {
	return FromContext(c).GetUser(id)
}

func GetUserByEmail(c context.Context, email string) (*model.User, error) {
	return FromContext(c).GetUserByEmail(email)
}

func GetUserByLogin(c context.Context, email string, password string) (*model.User, error) {
	return FromContext(c).GetUserByLogin(email, password)
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

// GetPost gets the user object corresponding to the given id.
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

func GetComment(c context.Context, id uint) (*model.Comment, error) {
	return FromContext(c).GetComment(id)
}

func CreateComment(c context.Context, comment *model.Comment) error {
	return FromContext(c).CreateComment(comment)
}

func UpdateComment(c context.Context, comment *model.Comment, id uint) (*model.Comment, error) {
	return FromContext(c).UpdateComment(comment, id)
}

func DeleteComment(c context.Context, id uint) error {
	return FromContext(c).DeleteComment(id)
}
