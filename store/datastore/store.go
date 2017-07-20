package datastore

import (
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ufukomer/tagon-api/store"

	"github.com/jinzhu/gorm"
)

// Datastore is an implementation of a model.Store built on top
// of the mysql driver with a relational database backend.
type Datastore struct {
	*gorm.DB
	Host     string
	DBName   string
	User     string
	Password string
	Port     int
}

type Option func(*Datastore)

func New(option Option) store.Store {
	d := &Datastore{}
	option(d)
	d.DB = open(d)
	return d
}

func open(d *Datastore) *gorm.DB {
	db, err := gorm.Open("mysql", d.User+":"+d.Password+"@tcp("+d.Host+":"+strconv.Itoa(d.Port)+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
