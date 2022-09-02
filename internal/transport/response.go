package transport

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
}

func NewResponse(c *gin.Context, status int, message string) {
	c.JSON(status, Response{
		Message: message,
	})
}
