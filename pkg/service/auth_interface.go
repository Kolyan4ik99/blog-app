package service

import (
	"net/http"
)

type AuthI interface {
	Signup(w http.ResponseWriter, r *http.Request)
	Signin(w http.ResponseWriter, r *http.Request)
}
