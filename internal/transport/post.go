package transport

import (
	"context"

	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/gin-gonic/gin"
)

type PostInterface interface {
	GetPosts(c *gin.Context)
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

func (p *Post) GetPosts(c *gin.Context) {

}

func (p *Post) GetPostByID(c *gin.Context) {

}

func (p *Post) UploadPost(c *gin.Context) {

}

func (p *Post) UpdatePostByID(c *gin.Context) {

}

func (p *Post) DeletePostByID(c *gin.Context) {

}
