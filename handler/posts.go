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

func PostPost(c *gin.Context) {
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := store.CreatePost(c, post); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, post)
}

func PatchPost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	in := new(model.Post)
	if err := c.Bind(in); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	post := &model.Post{
		Title:       in.Title,
		Content:     in.Content,
		UserID:      in.UserID,
		PublishedAt: in.PublishedAt,
		Vote:        in.Vote,
	}
	newPost, err := store.UpdatePost(c, post, uint(id))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, newPost)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := store.DeletePost(c, uint(id)); err != nil {
		c.String(http.StatusInternalServerError, "Error deleting post. %s", err)
		return
	}
	c.Status(http.StatusOK)
}
