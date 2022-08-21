package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Repo *repository.Users
	Ctx  *context.Context
}

func (a *Auth) Signup(c *gin.Context) {
	var user repository.UserInfo
	arr, err := io.ReadAll(c.Request.Body)
	if err != nil {
		BadRequest(c)
		return
	}
	defer c.Request.Body.Close()
	err = json.Unmarshal(arr, &user)
	if err != nil {
		BadRequest(c)
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	newUserId, err := a.Repo.Save(a.Ctx, &user)
	if err != nil {
		BadRequest(c)
		logger.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	logger.Logger.Infof("Sign-up was succesful completed, new user id = %d", newUserId)
	c.JSON(http.StatusOK, newUserId)
}

func (a *Auth) Signin(c *gin.Context) {
	// TODO необходимо реализовать аутентификацию
	panic("implement me")
}
