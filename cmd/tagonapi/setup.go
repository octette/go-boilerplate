package main

import (
	"github.com/ufukomer/tagon-api/store"
	"github.com/ufukomer/tagon-api/store/datastore"
	"github.com/urfave/cli"
)

func setupStore(c *cli.Context) store.Store {
	return datastore.New(datastore.Config{
		Host:     c.String("127.0.0.1"),
		DBName:   c.String("tagon-api"),
		DBUser:   c.String("root"),
		Password: c.String(""),
		Port:     c.Int("3306"),
	})
}
