package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
)

type AccessInterface interface {
	GetAccessPost(w http.ResponseWriter, r *http.Request)
	SetAccessPost(w http.ResponseWriter, r *http.Request)
}

type Access struct {
	ctx           context.Context
	accessService service.AccessInterface
}

func NewAccess(ctx context.Context, accessService service.AccessInterface) *Access {
	return &Access{ctx: ctx, accessService: accessService}
}

// GetAccessPost godoc
// @Summary      List post accesses
// @Description  Get all accesses in post by post_id.
// @Description  "is_all" - все ли пользователи могут получить доступ
// @Description  Остальные параметры определяют доступ для каждой записи отдельно
// @Security     ApiKeyAuth
// @Tags         access
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Success      200  {object}  []model.AccessOutput
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/access/{id} [get]
func (a *Access) GetAccessPost(w http.ResponseWriter, r *http.Request) {
	postId, err := parsePostId(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	accessPosts, err := a.accessService.GetAccessPost(a.ctx, postId)
	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponseBody(w, accessPosts)
}

// SetAccessPost godoc
// @Summary      Upload accesses
// @Description  Upload new accesses to posts by users
// @Security     ApiKeyAuth
// @Tags         access
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Param        input       body     model.AccessSetInput true "Body for new post"
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/access/{id} [post]
func (a *Access) SetAccessPost(w http.ResponseWriter, r *http.Request) {
	postId, err := parsePostId(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	accessPosts, err := parseAccess(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = a.accessService.SetAccessPost(a.ctx, postId, accessPosts)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	NewResponse(w, http.StatusCreated, "Created")
}

func parseAccess(r *http.Request) (*model.AccessSetInput, error) {
	var access model.AccessSetInput
	byteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteArr, &access)
	if err != nil {
		return nil, err
	}

	return &access, err
}
