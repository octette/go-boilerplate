package model

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Content     string
	UserID      uint
	PublishedAt string
	Vote        int32
}
