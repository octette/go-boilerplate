package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	
	"github.com/ufukomer/go-boilerplate/store"
)

func GetUserList(c *gin.Context) {
	users, err := store.GetUserList(c)
	if err != nil {
		fmt.Print(err)
		c.Error(err)
	}
	c.JSON(200, users)
}
