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

	api := e.Group("/api/users")
	{
		api.GET("", handler.GetUsers)
		api.GET("/:id", handler.GetUser)
		api.POST("", handler.PostUser)
		api.DELETE("/:id", handler.DeleteUser)
	}

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
