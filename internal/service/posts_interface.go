package service

import (
	"github.com/gin-gonic/gin"
)

type PostsI interface {
	GetPosts(c *gin.Context)
	GetPostByID(c *gin.Context)
	UploadPost(c *gin.Context)
	UpdatePostByID(c *gin.Context)
	DeletePostByID(c *gin.Context)
}
