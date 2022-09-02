package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

type Auth struct {
	ctx         context.Context
	authService service.AuthInterface
}

func NewAuth(ctx context.Context, authService service.AuthInterface) *Auth {
	return &Auth{ctx: ctx, authService: authService}
}

func (a *Auth) Signup(c *gin.Context) {
	var user model.UserInfo
	arr, err := io.ReadAll(c.Request.Body)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = json.Unmarshal(arr, &user)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	id, err := a.authService.SignUp(a.ctx, &user)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (a *Auth) Signin(c *gin.Context) {
	// TODO реализовать метод
	c.JSON(http.StatusInternalServerError, "")
}
