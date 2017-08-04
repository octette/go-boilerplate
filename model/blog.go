package model

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	UserID     uint
	Domain     string
	Authors    []User     `gorm:"many2many:blog_authors;"`
	Categories []Category `gorm:"many2many:blog_categories;"`
}
