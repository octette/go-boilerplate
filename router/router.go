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

	api.Use()
	{
		auth := api.Group("/auth")
		auth.Use(j.MiddlewareFunc())
		{
			auth.GET("/refresh_token", handler.Refresh)
		}

		users := api.Group("/users")
		{
			users.POST("", handler.PostUser)
			users.GET("", handler.GetUsers)
			users.GET("/:id", handler.GetUser)
			users.DELETE("/:id", handler.DeleteUser)
		}

		posts := api.Group("/posts")
		{
			posts.GET("", handler.GetPosts)
			posts.GET("/:id", handler.GetPost)
			posts.PATCH("/:id", handler.PatchPost)
			posts.POST("", handler.PostPost)
			posts.Use(j.MiddlewareFunc())
			{
				// posts.PATCH("/:id", handler.PatchPost)
				// posts.POST("", handler.PostPost)
				posts.DELETE("/:id", handler.DeletePost)
			}
		}
	}

	return e
}
