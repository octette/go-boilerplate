package main

import (
	"net/http"

	"github.com/ufukomer/tagon-api/router"
)

func main() {
	// setup the server and start the listener
	handler := router.Load()

	http.ListenAndServe(":8880", handler)
}
