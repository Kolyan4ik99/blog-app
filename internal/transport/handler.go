package transport

import (
	"github.com/gin-gonic/gin"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	authTransport AuthInterface
	postTransport PostInterface
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
				post.GET("/", h.postTransport.GetPostsByAuthor)
				post.GET("/:id", h.postTransport.GetPostByID)

				post.POST("/", h.postTransport.UploadPost)
				post.PUT("/:id", h.postTransport.UpdatePostByID)
				post.DELETE("/:id", h.postTransport.DeletePostByID)
			}
		}
	}
	return router
}

func NewHandler(auth AuthInterface, post PostInterface) *Handler {
	return &Handler{
		authTransport: auth,
		postTransport: post,
	}
}
