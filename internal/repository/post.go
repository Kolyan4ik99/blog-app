package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Kolyan4ik99/blog-app/internal/model"
)

type PostInterface interface {
	GetAll(ctx context.Context) ([]*model.PostInfo, error)
	GetById(ctx context.Context, id int64) (*model.PostInfo, error)
	Save(ctx context.Context, newPost *model.PostInfoInput) (int64, error)
	UpdateById(ctx context.Context, id int64, newPost *model.PostInfoUpdate) (*model.PostInfo, error)
	DeleteById(ctx context.Context, id int64) error
	GetAllPostTTLBefore(ctx context.Context, ttl time.Time) ([]*model.PostInfo, error)
}

type Post struct {
	maxId     int64
	postsInfo map[int64]*model.PostInfo
}

func NewPost() *Post {
	return &Post{
		maxId:     1,
		postsInfo: make(map[int64]*model.PostInfo, 50),
	}
}

func (p *Post) GetAll(ctx context.Context) ([]*model.PostInfo, error) {
	postsInfo := make([]*model.PostInfo, len(p.postsInfo))

	i := 0
	for _, info := range p.postsInfo {
		postsInfo[i] = info
		i++
	}
	return postsInfo, nil
}

func (p *Post) GetById(ctx context.Context, id int64) (*model.PostInfo, error) {
	post, exist := p.postsInfo[id]
	if !exist {
		return nil, sql.ErrNoRows
	}
	return post, nil
}

func (p *Post) Save(ctx context.Context, newPost *model.PostInfoInput) (int64, error) {
	valueTime, err := parseTime(newPost.TTL)
	if err != nil {
		return 0, err
	}

	p.postsInfo[p.maxId] = &model.PostInfo{
		Id:        p.maxId,
		Author:    newPost.Author,
		Header:    newPost.Header,
		Text:      newPost.Text,
		TTL:       valueTime,
		CreatedAt: time.Now(),
	}
	p.maxId++

	return p.maxId - 1, nil
}

func (p *Post) UpdateById(ctx context.Context, id int64, updatePost *model.PostInfoUpdate) (*model.PostInfo, error) {
	post, exist := p.postsInfo[id]
	if !exist {
		return nil, sql.ErrNoRows
	}

	valueTime, err := parseTime(updatePost.TTL)
	if err != nil {
		return nil, err
	}

	post.Text = updatePost.Text
	post.Header = updatePost.Header
	post.TTL = valueTime

	return post, nil
}

func (p *Post) DeleteById(ctx context.Context, id int64) error {
	_, exist := p.postsInfo[id]
	if !exist {
		return sql.ErrNoRows
	}

	delete(p.postsInfo, id)
	return nil
}

func (p *Post) GetAllPostTTLBefore(ctx context.Context, ttl time.Time) ([]*model.PostInfo, error) {
	var postsInfo []*model.PostInfo
	for _, val := range p.postsInfo {
		if val.TTL.Before(ttl) {
			postsInfo = append(postsInfo, val)
		}
	}

	return postsInfo, nil
}

func parseTime(timeFrom string) (*time.Time, error) {
	valueTime, err := time.Parse(time.RFC3339, timeFrom)
	if err != nil {
		return nil, err
	}
	return &valueTime, nil
}
