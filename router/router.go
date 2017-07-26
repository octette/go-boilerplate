package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ufukomer/go-boilerplate/handler"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(middleware...)

	api := e.Group("/api")
	{
		api.GET("users/:id", handler.GetUser)
		api.GET("users", handler.GetUsers)
	}

	return e
}
