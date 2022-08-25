package handler

import (
	"github.com/Kolyan4ik99/blog-app/internal/transport"
	"github.com/gin-gonic/gin"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	authTransport transport.AuthInterface
	postTransport transport.PostInterface
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sing-up", h.authTransport.Signup)
			auth.POST("/sing-in", h.authTransport.Signin)
		}
		api := v1.Group("/api")
		{
			post := api.Group("/post")
			{
				post.GET("/", h.postTransport.GetPosts)
				post.GET("/:id", h.postTransport.GetPostByID)

				post.POST("/", h.postTransport.UploadPost)
				post.PUT("/:id", h.postTransport.UpdatePostByID)
				post.DELETE("/:id", h.postTransport.DeletePostByID)
			}
		}
	}
	return router
}

func NewHandler(auth transport.AuthInterface, post transport.PostInterface) *Handler {
	return &Handler{
		authTransport: auth,
		postTransport: post,
	}
}
