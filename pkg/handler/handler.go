package handler

import (
	"github.com/Kolyan4ik99/blog-app/pkg/service"
	"github.com/gorilla/mux"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	AuthService service.AuthI
	PostService service.PostsI
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()

	authPath := "/auth"
	router.HandleFunc(authPath+"/sign-up", h.AuthService.Signup).Methods("POST")
	router.HandleFunc(authPath+"/sign-in", h.AuthService.Signin).Methods("POST")

	apiPostPath := "/api/post"
	router.HandleFunc(apiPostPath+"/", h.PostService.GetPosts).Methods("GET")
	router.HandleFunc(apiPostPath+"/{id:[0-9]+}", h.PostService.GetPostByID).Methods("GET")
	router.HandleFunc(apiPostPath+"/", h.PostService.UploadPost).Methods("POST")
	router.HandleFunc(apiPostPath+"/{id:[0-9]+}", h.PostService.UpdatePostByID).Methods("PUT")
	router.HandleFunc(apiPostPath+"/{id:[0-9]+}", h.PostService.DeletePostByID).Methods("DELETE")

	return router
}
