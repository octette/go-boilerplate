package main

import (
	"net/http"

	"github.com/ufukomer/tagon-api/router"
	"github.com/ufukomer/tagon-api/router/middleware"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "HOST",
		Name:   "host",
		Usage:  "host address",
		Value:  "localhost",
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

	// setup the server and start the listener
	handler := router.Load(
		middleware.Store(store),
	)

	http.ListenAndServe(":8880", handler)

	return nil
}
