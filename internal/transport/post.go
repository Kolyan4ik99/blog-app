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

var (
	ErrBadParam = errors.New("bad posts param - id")
)

type PostInterface interface {
	GetPostsByAuthor(c *gin.Context)
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

// GetPostsByAuthor godoc
// @Summary      list posts by author
// @Description  get list posts by author id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "author_id"
// @Success      200  {object}  model.PostInfo
// @Failure      400  {object}  transport.Response
// @Failure      404  {object}  transport.Response
// @Router       /api/post/{id} [get]
func (p *Post) GetPostsByAuthor(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	foundPosts, err := p.postService.GetPostsByAuthor(p.ctx, postId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			NewResponse(c, http.StatusNotFound,
				fmt.Sprintf("post with author_id = %d not found", postId))
		} else {
			NewResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, foundPosts)
}

// GetPostByID godoc
// @Summary      post by id
// @Description  get post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "author_id"
// @Success      200  {object}  model.PostInfo
// @Failure      400  {object}  transport.Response
// @Failure      404  {object}  transport.Response
// @Router       /api/post/{id} [get]
func (p *Post) GetPostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	foundPost, err := p.postService.GetPostByID(p.ctx, postId)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, foundPost)
}

// TODO add swag doc
func (p *Post) UploadPost(c *gin.Context) {
	post, err := parseBodyToPostInfo(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := p.postService.UploadPost(p.ctx, post)
	if err != nil {
		NewResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

// TODO add swag doc
func (p *Post) UpdatePostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		NewResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	post, err := parseBodyToPostInfo(c)
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

// TODO add swag doc
func (p *Post) DeletePostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
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

func (p *Post) parsePostId(c *gin.Context) (int64, error) {
	strId := c.Param("id")
	if strId == "" {
		return -1, ErrBadParam
	}
	num, err := strconv.Atoi(strId)
	if err != nil {
		return -1, ErrBadParam
	}

	return int64(num), nil
}

func parseBodyToPostInfo(c *gin.Context) (*model.PostInfo, error) {
	//bytes, err := io.ReadAll(c.Request.Body)
	//if err != nil {
	//	return nil, err
	//}
	var post model.PostInfo
	err := c.Bind(&post)
	//err = json.Unmarshal(bytes, &post)
	//if err != nil {
	//	return nil, err
	//}
	return &post, err
}
