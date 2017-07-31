package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ufukomer/go-boilerplate/model"
	"github.com/ufukomer/go-boilerplate/store"
)

func GetPosts(c *gin.Context) {
	posts, err := store.GetPostList(c)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting post list. %s", err)
		return
	}
	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	post, err := store.GetPost(c, uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Cannot find post. %s", err)
		return
	}
	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	in := &model.Post{}
	err := c.Bind(in)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	post := &model.Post{
		Title:       in.Title,
		Content:     in.Content,
		UserID:      in.UserID,
		PublishedAt: in.PublishedAt,
	}
	if err = store.CreatePost(c, post); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	in := &model.Post{}
	err := c.Bind(in)
	post := &model.Post{
		Title:       in.Title,
		Content:     in.Content,
		UserID:      in.UserID,
		PublishedAt: in.PublishedAt,
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if post, err = store.UpdatePost(c, post, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := store.DeletePost(c, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting post. %s", err)
		return
	}
	c.Status(http.StatusOK)
}
