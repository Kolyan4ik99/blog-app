package transport

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler В инициализации использую интерфейсы, дабы при изменении
// реализации сервисов не изменять реализацию InitRouter
type Handler struct {
	authTransport   AuthInterface
	postTransport   PostInterface
	accessTransport AccessInterface
}

func NewHandler(authTransport AuthInterface, postTransport PostInterface, accessTransport AccessInterface) *Handler {
	return &Handler{authTransport: authTransport, postTransport: postTransport, accessTransport: accessTransport}
}

// InitRouter конструктор роута с эндпоинтами
func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sing-up", h.authTransport.Signup)
			auth.POST("/sing-in", h.authTransport.Signin)
		}
		api := v1.Group("/api")
		{
			api.Use(h.authMiddleware)
			post := api.Group("/post")
			{
				post.GET("/", h.postTransport.GetAllPosts)
				post.GET("/:id", h.postTransport.GetPostByID)

				post.POST("/", h.postTransport.UploadPost)
				post.PUT("/:id", h.postTransport.UpdatePostByID)
				post.DELETE("/:id", h.postTransport.DeletePostByID)

				access := post.Group("/access")
				{
					access.GET("/:id", h.accessTransport.GetAccessPost)
					access.POST("/:id", h.accessTransport.SetAccessPost)
				}
			}
		}
	}
	return router
}
