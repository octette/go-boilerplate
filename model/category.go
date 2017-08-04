package model

import "github.com/jinzhu/gorm"

// CategoryType represents categories in db.
type CategoryType string

const (
	Science CategoryType = "science"
	Nature               = "nature"
	Earth                = "earth"
)

type Category struct {
	gorm.Model
	Name CategoryType
}
