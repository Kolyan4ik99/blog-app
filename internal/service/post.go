package service

import (
	"context"
	"errors"
	"time"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type PostInterface interface {
	GetAllPosts(ctx context.Context) ([]*model.PostInfo, error)
	GetPostByID(ctx context.Context, id int64) (*model.PostInfo, error)
	UploadPost(ctx context.Context, newPost *model.PostInfoInput) (int64, error)
	UpdatePostByID(ctx context.Context, id int64, newPost *model.PostInfoUpdate) (*model.PostInfo, error)
	DeletePostByID(ctx context.Context, id int64) error
}

var (
	ErrBadTTL = errors.New("bad ttl param")
)

type Post struct {
	repo repository.PostInterface
}

func NewPost(repo repository.PostInterface) *Post {
	return &Post{repo: repo}
}

func (p *Post) GetAllPosts(ctx context.Context) ([]*model.PostInfo, error) {
	return p.repo.GetAll(ctx)
}

func (p *Post) GetPostByID(ctx context.Context, id int64) (*model.PostInfo, error) {
	return p.repo.GetById(ctx, id)
}

func (p *Post) UploadPost(ctx context.Context, newPost *model.PostInfoInput) (int64, error) {
	parseTime, err := time.Parse(time.RFC3339, newPost.TTL)
	if err != nil {
		return 0, err
	}
	if parseTime.Before(time.Now()) {
		return 0, ErrBadTTL
	}
	return p.repo.Save(ctx, newPost)
}

func (p *Post) UpdatePostByID(ctx context.Context, id int64, newPost *model.PostInfoUpdate) (*model.PostInfo, error) {
	parseTime, err := time.Parse(time.RFC3339, newPost.TTL)
	if err != nil {
		return nil, err
	}
	if parseTime.Before(time.Now()) {
		err = p.repo.DeleteById(ctx, id)
		return nil, err
	}
	return p.repo.UpdateById(ctx, id, newPost)
}

func (p *Post) DeletePostByID(ctx context.Context, id int64) error {
	return p.repo.DeleteById(ctx, id)
}
