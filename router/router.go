package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/handler"
	mw "github.com/ufukomer/go-boilerplate/router/middleware"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(middleware...)

	authMiddleware := mw.JWT()

	e.POST("/login", authMiddleware.LoginHandler)

	api := e.Group("/api")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		auth := api.Group("/auth")
		{
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		}

		users := api.Group("/users")
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUser)
		users.POST("", handler.PostUser)
		users.DELETE("/:id", handler.DeleteUser)
	}

	return e
}
