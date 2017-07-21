package datastore

import "github.com/ufukomer/tagon-api/model"

func (db *Datastore) GetUserList() ([]*model.User, error) {
	var users = []*model.User{}
	var err = db.Find(&users).Error

	return users, err
}
func (db *Datastore) GetUser(ID string) (*model.User, error) {
	var user = *model.User{}
	var err = db.First(user, ID)

	return user
}
