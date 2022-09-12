package transport

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

type PostInterface interface {
	GetAllPosts(c *gin.Context)
	GetPostByID(c *gin.Context)
	UploadPost(c *gin.Context)
	UpdatePostByID(c *gin.Context)
	DeletePostByID(c *gin.Context)
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
func (p *Post) GetAllPosts(c *gin.Context) {
	foundPosts, err := p.postService.GetAllPosts(p.ctx)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundPosts)
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
func (p *Post) GetPostByID(c *gin.Context) {
	postId, err := parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	foundPost, err := p.postService.GetPostByID(p.ctx, postId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewResponse(c, http.StatusNotFound, err.Error())
		} else {
			NewResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, foundPost)
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
func (p *Post) UploadPost(c *gin.Context) {
	post, err := parseBodyToInput(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := p.postService.UploadPost(p.ctx, post)
	if err != nil {
		if errors.Is(err, service.ErrBadTTL) {
			NewResponse(c, http.StatusBadRequest, err.Error())
		} else {
			NewResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	NewResponse(c, http.StatusCreated, fmt.Sprintf("post_id: %d", id))
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
func (p *Post) UpdatePostByID(c *gin.Context) {
	postId, err := parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	post, err := parseBodyToUpdate(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatePost, err := p.postService.UpdatePostByID(p.ctx, postId, post)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updatePost)
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
func (p *Post) DeletePostByID(c *gin.Context) {
	postId, err := parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = p.postService.DeletePostByID(p.ctx, postId)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func parsePostId(c *gin.Context) (int64, error) {
	strId := c.Param("id")
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

func parseBodyToInput(c *gin.Context) (*model.PostInfoInput, error) {
	var post model.PostInfoInput
	err := c.Bind(&post)
	return &post, err
}

func parseBodyToUpdate(c *gin.Context) (*model.PostInfoUpdate, error) {
	var post model.PostInfoUpdate
	err := c.Bind(&post)
	return &post, err
}
