package datastore

import "github.com/ufukomer/go-boilerplate/model"

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
