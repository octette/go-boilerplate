package datastore

import (
	"github.com/ufukomer/go-boilerplate/model"
)

func (db *Datastore) GetComment(id uint) (*model.Comment, error) {
	var comment = new(model.Comment)
	var err = db.First(&comment, id).Error

	return comment, err
}

func (db *Datastore) CreateComment(comment *model.Comment) error {
	var err = db.Create(&comment).Error
	return err
}

func (db *Datastore) UpdateComment(comment *model.Comment, id uint) (*model.Comment, error) {
	newComment := new(model.Comment)
	if err := db.First(newComment, id).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&newComment).Updates(comment).Error; err != nil {
		return nil, err
	}

	return newComment, nil
}

func (db *Datastore) DeleteComment(id uint) error {
	var comment = new(model.Comment)
	if err := db.Find(&comment, id).Error; err != nil {
		return err
	}
	var err = db.Delete(&comment).Error

	return err
}
