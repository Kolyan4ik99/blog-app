package service

import (
	"context"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type PostInterface interface {
	GetPostsByAuthor(ctx context.Context, id int64) ([]*model.PostInfo, error)
	GetPostByID(ctx context.Context, id int64) (*model.PostInfo, error)
	UploadPost(ctx context.Context, newPost *model.PostInfo) (int64, error)
	UpdatePostByID(ctx context.Context, id int64, newPost *model.PostInfo) (*model.PostInfo, error)
	DeletePostByID(ctx context.Context, id int64) error
}

type Post struct {
	repo repository.PostInterface
}

func NewPost(repo repository.PostInterface) *Post {
	return &Post{repo: repo}
}

func (p *Post) GetPostsByAuthor(ctx context.Context, id int64) ([]*model.PostInfo, error) {
	return p.repo.GetAllByAuthorId(ctx, id)
}

func (p *Post) GetPostByID(ctx context.Context, id int64) (*model.PostInfo, error) {
	return p.repo.GetById(ctx, id)
}

func (p *Post) UploadPost(ctx context.Context, newPost *model.PostInfo) (int64, error) {
	return p.repo.Save(ctx, newPost)
}

func (p *Post) UpdatePostByID(ctx context.Context, id int64, newPost *model.PostInfo) (*model.PostInfo, error) {
	return p.repo.UpdateById(ctx, id, newPost)
}

func (p *Post) DeletePostByID(ctx context.Context, id int64) error {
	return p.repo.DeleteById(ctx, id)
}
