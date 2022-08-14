package handler

import (
	"github.com/Kolyan4ik99/blog-app/pkg/service"
	"github.com/gin-gonic/gin"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	AuthService service.AuthI
	PostService service.PostsI
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.AuthService.Signin)
		auth.POST("/sign-up", h.AuthService.Signup)
	}

	api := router.Group("/api")
	{
		postList := api.Group("/posts")
		{
			// Получение всех постов
			postList.GET("/", h.PostService.GetPosts)
		}

		post := api.Group("/post")
		{
			// Получение поста по id
			post.GET("/:id", h.PostService.GetPostByID)
			post.POST("/", h.PostService.UploadPost)
			post.PUT("/:id", h.PostService.UpdatePostByID)
			post.DELETE("/:id", h.PostService.DeletePostByID)
		}
	}

	return router
}
