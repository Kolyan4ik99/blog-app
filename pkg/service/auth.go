package service

import (
	"context"
	"encoding/json"
	"github.com/Kolyan4ik99/blog-app/internal"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Auth struct {
	Repo *repository.Users
	Ctx  *context.Context
}

func (a *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var user repository.UserInfo
	arr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		BadRequest(w)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(arr, &user)
	if err != nil {
		BadRequest(w)
		internal.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	newUser, err := a.Repo.Save(a.Ctx, &user)
	if err != nil {
		BadRequest(w)
		internal.Logger.Warningf("Bad request: %s\n", err)
		return
	}

	internal.Logger.Infof("Sign-up was succesful completed %v", newUser)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.FormatInt(newUser.Id, 10)))
}

func (a *Auth) Signin(w http.ResponseWriter, r *http.Request) {
	// TODO необходимо реализовать аутентификацию
	panic("implement me")
}
