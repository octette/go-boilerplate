package datastore

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/ufukomer/go-boilerplate/model"
)

func (db *Datastore) GetUserList() ([]*model.User, error) {
	var users = []*model.User{}
	var err = db.Find(&users).Error

	return users, err
}

func (db *Datastore) GetUser(id uint) (*model.User, error) {
	var user = &model.User{}
	var err = db.First(&user, id).Error

	return user, err
}

func (db *Datastore) GetUserByEmail(email string) (*model.User, error) {
	var user = &model.User{}
	var err = db.Where(&model.User{Email: email}).First(&user).Error

	return user, err
}

func (db *Datastore) GetUserByLogin(email string, password string) (*model.User, error) {
	var user = &model.User{}

	if err := db.Where(&model.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Hash, []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (db *Datastore) CreateUser(user *model.User) error {
	var err = db.Create(&user).Error

	return err
}

func (db *Datastore) DeleteUser(id uint) error {
	var user = &model.User{}
	if err := db.Find(&user, id).Error; err != nil {
		return err
	}
	var err = db.Delete(&user).Error

	return err
}
