package service

import (
	"github.com/gin-gonic/gin"
)

type AuthI interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}
