package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/store"
)

func GetUserList(c *gin.Context) {
	users := store.GetUserList(c)
	c.JSON(200, users)
}

// func GetUser(c *gin.Context)  {
//
// }
