package service

import (
	"net/http"
)

type PostsI interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	UploadPost(w http.ResponseWriter, r *http.Request)
	UpdatePostByID(w http.ResponseWriter, r *http.Request)
	DeletePostByID(w http.ResponseWriter, r *http.Request)
}
