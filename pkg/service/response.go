package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response обработчик сообщений об ошибках
type Response struct {
	Message string `json:"message,required"`
	Status  int    `json:"status,required"`
}

func BadRequest(c *gin.Context, msg string) {
	if msg == "" {
		msg = "Bad request"
	}
	resp := Response{
		Message: msg,
		Status:  http.StatusBadRequest,
	}
	c.JSON(resp.Status, resp)
}

func NotFound(c *gin.Context, msg string) {
	if msg == "" {
		msg = "Not found"
	}
	resp := Response{
		Message: msg,
		Status:  http.StatusNotFound,
	}
	c.JSON(resp.Status, resp)
}

func InternalServerError(c *gin.Context, msg string) {
	if msg == "" {
		msg = "Something wrong"
	}
	resp := Response{
		Message: msg,
		Status:  http.StatusInternalServerError,
	}
	c.JSON(resp.Status, resp)
}
