package main

import (
	"github.com/ufukomer/tagon-api/store"
	"github.com/ufukomer/tagon-api/store/datastore"
	"github.com/urfave/cli"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(datastore.Config{
		Host:     c.String("host"),
		DBName:   c.String("mysql-dbname"),
		DBUser:   c.String("mysql-dbuser"),
		Password: c.String("mysql-password"),
		Port:     c.Int("mysql-port"),
	})
}
