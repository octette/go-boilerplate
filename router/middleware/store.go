package middleware

import (
	"github.com/gin-gonic/gin"
	
	"github.com/ufukomer/go-boilerplate/store"
)

// Store is a middleware function that initializes the Datastore and attaches to
// the context of every http.Request.
func Store(v store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, v)
		c.Next()
	}
}
