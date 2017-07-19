package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ufukomer/tagon-api/handler"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(middleware...)

	v3 := e.Group("/v3")
	{
		v3.GET("users", handler.GetUserList)
	}

	return e
}
