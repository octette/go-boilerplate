package model

import "github.com/jinzhu/gorm"

// Comment represents posts' comments.
type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	PostID  uint
	Vote    uint32
}
