package datastore

import "github.com/ufukomer/go-boilerplate/model"

func (db *Datastore) GetUserList() ([]*model.User, error) {
	var users = []*model.User{}
	var err = db.Find(&users).Error

	return users, err
}

func (db *Datastore) GetUser(id int64) (*model.User, error) {
	var user = &model.User{}
	var err = db.First(&user, id).Error

	return user, err
}
