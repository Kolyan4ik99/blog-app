package handler

import (
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	authService service.AuthI
	postService service.PostsI
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sing-up", h.authService.Signup)
			auth.POST("/sing-in", h.authService.Signin)
		}
		api := v1.Group("/api")
		{
			post := api.Group("/post")
			{
				post.GET("/", h.postService.GetPosts)
				post.GET("/:id", h.postService.GetPostByID)

				post.POST("/", h.postService.UploadPost)
				post.PUT("/:id", h.postService.UpdatePostByID)
				post.DELETE("/:id", h.postService.DeletePostByID)
			}
		}
	}
	return router
}

func NewHandler(auth service.AuthI, post service.PostsI) *Handler {
	return &Handler{
		authService: auth,
		postService: post,
	}
}
