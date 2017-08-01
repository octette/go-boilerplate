package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/handler"
	j "github.com/ufukomer/go-boilerplate/router/middleware/jwt"
)

// Load returns a handler with routes specified with given middlewares.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(middleware...)

	e.POST("/login", handler.Login)

	api := e.Group("/api")
	users := api.Group("/users")
	api.Use(j.JWT())
	{
		auth := api.Group("/auth")
		{
			auth.GET("/refresh_token", handler.RefreshHandler)
		}

		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUser)
		users.DELETE("/:id", handler.DeleteUser)
	}

	users.POST("", handler.PostUser)

	api = e.Group("/api/posts")
	{
		api.GET("", handler.GetPosts)
		api.GET("/:id", handler.GetPost)
		api.PUT("/:id", handler.UpdatePost)
		api.POST("", handler.CreatePost)
		api.DELETE("/:id", handler.DeletePost)
	}

	return e
}
