package datastore

import "github.com/ufukomer/tagon-api/model"

func (db *datastore) GetUserList() ([]*model.User, error) {
	var users = []*model.User{}
	var err = db.Find(&users).Error

	return users, err
}
