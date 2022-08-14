package service

import "github.com/gin-gonic/gin"

type AuthI interface {
	Signup(ctx *gin.Context)
	Signin(ctx *gin.Context)
}
