package datastore

import (
	"github.com/ufukomer/go-boilerplate/model"
)

func (db *Datastore) GetPostList() ([]*model.Post, error) {
	var posts = []*model.Post{}
	var err = db.Find(&posts).Error
	return posts, err
}

func (db *Datastore) GetPost(id uint) (*model.Post, error) {
	var post = &model.Post{}
	var err = db.First(&post, id).Error

	return post, err
}

func (db *Datastore) CreatePost(post *model.Post) error {
	var err = db.Create(&post).Error
	return err
}

func (db *Datastore) UpdatePost(post *model.Post, id uint) (*model.Post, error) {
	newPost := new(model.Post)
	if err := db.First(newPost, id).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&newPost).Updates(post).Error; err != nil {
		return nil, err
	}

	return newPost, nil
}

func (db *Datastore) DeletePost(id uint) error {
	var post = &model.Post{}
	if err := db.Find(&post, id).Error; err != nil {
		return err
	}
	var err = db.Delete(&post).Error

	return err
}
