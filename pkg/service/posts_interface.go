package service

import "github.com/gin-gonic/gin"

type PostsI interface {
	GetPosts(ctx *gin.Context)
	GetPostByID(ctx *gin.Context)
	UploadPost(ctx *gin.Context)
	UpdatePostByID(ctx *gin.Context)
	DeletePostByID(ctx *gin.Context)
}
