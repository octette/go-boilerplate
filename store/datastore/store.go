package datastore

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ufukomer/tagon-api/store"

	"github.com/jinzhu/gorm"
)

// datastore is an implementation of a model.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*gorm.DB
}

type Config struct {
	Host     string
	DBName   string
	DBUser   string
	Password string
	Port     int
}

func New(config Config) store.Store {
	return &datastore{
		DB: open(config),
	}

}

func open(config Config) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
