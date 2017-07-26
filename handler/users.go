package handler

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/store"
)

func GetUsers(c *gin.Context) {
	users, err := store.GetUserList(c)
	if err != nil {
		c.String(500, "Error getting user list. %s", err)
		return
	}
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.String(404, "Cannot get user with id '%v'.", c.Param("id"))
		logrus.Errorln(err)
		return
	}

	user, err := store.GetUser(c, uint(id))
	if err != nil {
		c.String(404, "Cannot find user. %s", err)
		return
	}
	c.JSON(200, user)
}
