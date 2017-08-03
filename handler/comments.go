package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ufukomer/go-boilerplate/model"
	"github.com/ufukomer/go-boilerplate/store"
)

func GetComment(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	comment, err := store.GetComment(c, uint(id))
	if err != nil {
		c.String(http.StatusInternalServerError, "Cannot find comment. %s", err)
		return
	}
	c.JSON(http.StatusOK, comment)
}

func PostComment(c *gin.Context) {
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	if err := store.CreateComment(c, comment); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
}

func PathComment(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	in := new(model.Comment)

	if err := c.Bind(in); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	comment := &model.Comment{
		Content: in.Content,
		UserID:  in.UserID,
		PostID:  in.PostID,
		Vote:    in.Vote,
	}
	if _, err := store.UpdateComment(c, comment, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := store.DeleteComment(c, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting comment. %s", err)
		return
	}
	c.Status(http.StatusOK)
}
