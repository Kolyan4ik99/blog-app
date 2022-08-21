package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(c *gin.Context) {
	c.Status(http.StatusBadRequest)
}

func NotFound(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func InternalServerError(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
}
