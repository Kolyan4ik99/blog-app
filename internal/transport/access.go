package transport

import (
	"context"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

type AccessInterface interface {
	GetAccessPost(ctx *gin.Context)
	SetAccessPost(ctx *gin.Context)
}

type Access struct {
	ctx           context.Context
	accessService service.AccessInterface
}

func NewAccess(ctx context.Context, accessService service.AccessInterface) *Access {
	return &Access{ctx: ctx, accessService: accessService}
}

// GetAccessPost godoc
// @Summary      list accesses by post_id
// @Description  get all accesses int post
// @Security     ApiKeyAuth
// @Tags         access
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Success      200  {object}  []model.AccessOutput
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/access/{id} [get]
func (a *Access) GetAccessPost(c *gin.Context) {
	postId, err := parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	accessPosts, err := a.accessService.GetAccessPost(a.ctx, postId)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, accessPosts)
}

// SetAccessPost godoc
// @Summary      list accesses by post_id
// @Description  get all accesses int post
// @Security     ApiKeyAuth
// @Tags         access
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Param        input       body     model.AccessSetInput true "Body for new post"
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/access/{id} [post]
func (a *Access) SetAccessPost(c *gin.Context) {
	postId, err := parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var accessPosts model.AccessSetInput
	err = c.Bind(&accessPosts)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = a.accessService.SetAccessPost(a.ctx, postId, &accessPosts)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	NewResponse(c, http.StatusCreated, "Created")
}
