package transport

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
)

type PostInterface interface {
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	GetPostByID(w http.ResponseWriter, r *http.Request)
	UploadPost(w http.ResponseWriter, r *http.Request)
	UpdatePostByID(w http.ResponseWriter, r *http.Request)
	DeletePostByID(w http.ResponseWriter, r *http.Request)
}

type Post struct {
	ctx         context.Context
	postService service.PostInterface
}

func NewPost(ctx context.Context, newPostService service.PostInterface) *Post {
	return &Post{
		ctx:         ctx,
		postService: newPostService,
	}
}

var (
	ErrBadParam = errors.New("bad posts param - id")
)

// GetAllPosts godoc
// @Summary      List all posts
// @Description  Get all posts in list
// @Security     ApiKeyAuth
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.PostInfo
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/ [get]
func (p *Post) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	foundPosts, err := p.postService.GetAllPosts(p.ctx)
	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponseBody(w, foundPosts)
}

// GetPostByID godoc
// @Summary      Get post
// @Description  Get post by post_id
// @Security     ApiKeyAuth
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Success      200  {object}  model.PostInfo
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/{id} [get]
func (p *Post) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postId, err := parsePostId(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	foundPost, err := p.postService.GetPostByID(p.ctx, postId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewResponse(w, http.StatusNotFound, err.Error())
		} else {
			NewResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	writeResponseBody(w, foundPost)
}

// UploadPost godoc
// @Summary      Upload new post
// @Description  Method create new post by required body
// @Security     ApiKeyAuth
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        input       body     model.PostInfoInput true "Body for new post"
// @Success      200  {object}  transport.Response
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/ [post]
func (p *Post) UploadPost(w http.ResponseWriter, r *http.Request) {
	post, err := parseBodyToInput(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := p.postService.UploadPost(p.ctx, post)
	if err != nil {
		if errors.Is(err, service.ErrBadTTL) {
			NewResponse(w, http.StatusBadRequest, err.Error())
		} else {
			NewResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	NewResponse(w, http.StatusCreated, fmt.Sprintf("post_id: %d", id))
}

// UpdatePostByID godoc
// @Summary      Update post
// @Description  Method update post by required body
// @Security     ApiKeyAuth
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Param        input       body     model.PostInfoUpdate true "Body for update"
// @Success      200  {object}  model.PostInfo
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/{id} [put]
func (p *Post) UpdatePostByID(w http.ResponseWriter, r *http.Request) {
	postId, err := parsePostId(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	post, err := parseBodyToUpdate(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	updatePost, err := p.postService.UpdatePostByID(p.ctx, postId, post)
	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponseBody(w, updatePost)
}

// DeletePostByID godoc
// @Summary      Delete post
// @Description  Method delete post by post_id
// @Security     ApiKeyAuth
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "post_id"
// @Success      200  {object}  transport.Response
// @Failure      400,401,404,500 {object} transport.Response
// @Router       /api/post/{id} [delete]
func (p *Post) DeletePostByID(w http.ResponseWriter, r *http.Request) {
	postId, err := parsePostId(r)
	if err != nil {
		NewResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = p.postService.DeletePostByID(p.ctx, postId)
	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

func parsePostId(r *http.Request) (int64, error) {
	strs := strings.Split(r.URL.RequestURI(), "/")
	strId := strs[len(strs)-1]
	if strId == "" {
		return -1, ErrBadParam
	}
	num, err := strconv.Atoi(strId)
	if err != nil {
		return -1, ErrBadParam
	}
	if num < 0 {
		return -1, ErrBadParam
	}

	return int64(num), nil
}

func parseBodyToInput(r *http.Request) (*model.PostInfoInput, error) {
	var post model.PostInfoInput
	byteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteArr, &post)
	if err != nil {
		return nil, err
	}

	return &post, err
}

func parseBodyToUpdate(r *http.Request) (*model.PostInfoUpdate, error) {
	var post model.PostInfoUpdate
	byteArr, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteArr, &post)
	if err != nil {
		return nil, err
	}

	return &post, err
}

func writeResponseBody(w http.ResponseWriter, body interface{}) {
	byteArr, err := json.Marshal(body)
	if err != nil {
		NewResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteArr)
}
