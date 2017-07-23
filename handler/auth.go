package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ufukomer/go-boilerplate/store"
	"github.com/ufukomer/go-boilerplate/util"
)

type Login struct {
	Email    string `email:"json"`
	Password string `password:"json"`
}

func PostLogin(c *gin.Context) {
	var login Login
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println(err)
	}
	user, err := store.GetUserByEmail(c, login.Email)
	if err != nil {
		fmt.Println(err)
	}
	comparePassword := util.CheckPasswordHash(login.Password, user.Password)
	fmt.Println(comparePassword)
	if login.Email == user.Email && comparePassword == true {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}
