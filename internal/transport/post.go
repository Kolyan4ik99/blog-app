package transport

import (
	"context"
	"errors"
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

func (p *Post) GetPostsByAuthor(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		BadRequest(c)
		return
	}

	foundPosts, err := p.postService.GetPostsByAuthor(p.ctx, postId)
	if err != nil {
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, foundPosts)
}

func (p *Post) GetPostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		BadRequest(c)
		return
	}

	foundPost, err := p.postService.GetPostByID(p.ctx, postId)
	if err != nil {
		InternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, foundPost)
}

func (p *Post) UploadPost(c *gin.Context) {
	post, err := parseBodyToPostInfo(c)
	if err != nil {
		BadRequest(c)
		return
	}

	id, err := p.postService.UploadPost(p.ctx, post)
	if err != nil {
		InternalServerError(c)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (p *Post) UpdatePostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		BadRequest(c)
		return
	}
	post, err := parseBodyToPostInfo(c)
	if err != nil {
		BadRequest(c)
		return
	}

	updatePost, err := p.postService.UpdatePostByID(p.ctx, postId, post)
	if err != nil {
		InternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, updatePost)
}

func (p *Post) DeletePostByID(c *gin.Context) {
	postId, err := p.parsePostId(c)
	if err != nil {
		BadRequest(c)
		return
	}
	err = p.postService.DeletePostByID(p.ctx, postId)
	if err != nil {
		InternalServerError(c)
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
