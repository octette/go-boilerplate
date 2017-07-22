package main

import (
	"github.com/ufukomer/go-boilerplate/store"
	"github.com/ufukomer/go-boilerplate/store/datastore"
	"github.com/urfave/cli"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(func(d *datastore.Datastore) {
		d.Host = c.String("host")
		d.DBName = c.String("mysql-dbname")
		d.User = c.String("mysql-dbuser")
		d.Password = c.String("mysql-password")
		d.Port = c.Int("mysql-port")
	})
}
