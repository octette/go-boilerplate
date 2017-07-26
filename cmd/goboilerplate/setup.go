package main

import (
	"github.com/urfave/cli"

	"github.com/ufukomer/go-boilerplate/store/datastore"
)

func setupStore(c *cli.Context) *datastore.Datastore {
	return datastore.New(func(d *datastore.Datastore) {
		d.Host = c.String("host")
		d.DBName = c.String("mysql-dbname")
		d.User = c.String("mysql-dbuser")
		d.Password = c.String("mysql-password")
		d.Port = c.Int("mysql-port")
	})
}
