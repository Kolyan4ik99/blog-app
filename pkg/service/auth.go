package service

import (
	"encoding/json"
	"fmt"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Auth struct {
	Repo repository.Users
}

func (a *Auth) Signup(c *gin.Context) {
	var user repository.UserInfo
	arr, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		BadRequest(c, "")
		return
	}
	err = json.Unmarshal(arr, &user)
	fmt.Println(user)
	if err != nil {
		BadRequest(c, "")
		return
	}
	err = a.Repo.Save(&user)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusCreated, fmt.Sprintf("User was successful created"))
}

func (a *Auth) Signin(c *gin.Context) {
	// TODO необходимо реализовать аутентификацию
	panic("implement me")
}
