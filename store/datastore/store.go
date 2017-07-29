package datastore

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ufukomer/go-boilerplate/model"
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

type options func(*Datastore)

func New(opts options) *Datastore {
	d := &Datastore{}
	opts(d)
	d.DB = open(d)
	return d
}

func open(d *Datastore) *gorm.DB {
	db, err := gorm.Open("mysql", d.User+":"+d.Password+"@tcp("+d.Host+":"+strconv.Itoa(d.Port)+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})

	return db
}
