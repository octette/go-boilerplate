package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/model"
	"github.com/ufukomer/go-boilerplate/store"
)

func GetUsers(c *gin.Context) {
	users, err := store.GetUserList(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting user list. %s", err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	user, err := store.GetUser(c, uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Cannot find user. %s", err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	in := &model.User{}
	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}
	if err = user.Validate(); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err = store.CreateUser(c, user); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := store.DeleteUser(c, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting user. %s", err)
		return
	}
	c.Status(http.StatusOK)
}
