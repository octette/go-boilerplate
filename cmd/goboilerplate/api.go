package main

import (
	"net/http"

	"github.com/urfave/cli"

	"github.com/ufukomer/go-boilerplate/router"
	"github.com/ufukomer/go-boilerplate/router/middleware"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "HOST",
		Name:   "host",
		Usage:  "host address",
		Value:  "localhost",
	},
	cli.StringFlag{
		EnvVar: "PORT",
		Name:   "port",
		Usage:  "port number",
		Value:  "8880",
	},
	cli.StringFlag{
		EnvVar: "MYSQL_DBNAME",
		Name:   "mysql-dbname",
		Usage:  "msql database name",
		Value:  "go-boilerplate",
	},
	cli.StringFlag{
		EnvVar: "MYSQL_DBUSER",
		Name:   "mysql-dbuser",
		Usage:  "mysql database user",
		Value:  "root",
	},
	cli.StringFlag{
		EnvVar: "MYSQL_PASSWORD",
		Name:   "mysql-password",
		Usage:  "mysql password",
		Value:  "",
	},
	cli.IntFlag{
		EnvVar: "MYSQL_PORT",
		Name:   "mysql-port",
		Usage:  "mysql port",
		Value:  3306,
	},
}

func api(c *cli.Context) error {

	store := setupStore(c)
	// close database
	defer store.Close()

	// setup the server and start the listener
	handler := router.Load(
		middleware.Store(store),
	)

	http.ListenAndServe(":"+c.String("port"), handler)

	return nil
}
