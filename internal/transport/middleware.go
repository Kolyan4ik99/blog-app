package transport

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrEmptyToken   = errors.New("invalid token")
	ErrInvalidToken = errors.New("invalid token")
)

func (h *Handler) authMiddleware(c *gin.Context) {
	token, err := parseToken(c)
	if err != nil {
		NewResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	if !h.authTransport.CheckToken(token) {
		NewResponse(c, http.StatusUnauthorized, ErrInvalidToken.Error())
		c.Abort()
		return
	}
	c.Next()
}

func parseToken(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")
	if header == "" {
		return "", ErrEmptyToken
	}
	headerAuth := strings.Split(header, " ")
	if len(headerAuth) != 2 || headerAuth[0] != "Bearer" {
		return "", ErrInvalidToken
	}
	return headerAuth[1], nil
}
