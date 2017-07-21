package datastore

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ufukomer/tagon-api/model"
	"github.com/ufukomer/tagon-api/store"
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

func New(opts options) store.Store {
	d := &Datastore{}
	opts(d)
	d.DB = open(d)
	return d
}

func open(d *Datastore) *gorm.DB {
	db, err := gorm.OpOpen("mysql", d.User+":"+d.Password+"@tcp("+d.Host+":"+strconv.Itoa(d.Port)+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database connection failed")
	}
	defer db.Close()

	db.AutoMigrate(&model.User{})

	return db
}
