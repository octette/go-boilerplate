package main

import (
	"net/http"

	"github.com/ufukomer/tagon-api/router"
	"github.com/ufukomer/tagon-api/router/middleware"
	"github.com/ufukomer/tagon-api/store/datastore"
)

func main() {

	db := datastore.New(datastore.Config{
		Host:     "localhost",
		DBName:   "tagon-api",
		DBUser:   "root",
		Password: "",
		Port:     3306,
	})
	// setup the server and start the listener
	handler := router.Load(middleware.Store(db))

	http.ListenAndServe(":8880", handler)
}
