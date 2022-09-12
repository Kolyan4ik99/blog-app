package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type PostInterface interface {
	GetAllByAuthorId(ctx context.Context) ([]*model.PostInfo, error)
	GetById(ctx context.Context, id int64) (*model.PostInfo, error)
	Save(ctx context.Context, newPost *model.PostInfoInput) (int64, error)
	UpdateById(ctx context.Context, id int64, newPost *model.PostInfoUpdate) (*model.PostInfo, error)
	DeleteById(ctx context.Context, id int64) error
	GetAllPostTTLBefore(ctx context.Context, ttl time.Time) ([]*model.PostInfo, error)
}

type Post struct {
	con *sqlx.DB
}

func NewPost(con *sqlx.DB) *Post {
	return &Post{con: con}
}

func (p *Post) GetAllByAuthorId(ctx context.Context) ([]*model.PostInfo, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, postgres.PostTable)

	var postsInfo []*model.PostInfo
	err := p.con.SelectContext(ctx, &postsInfo, query)
	if err != nil {
		return nil, err
	}
	return postsInfo, nil
}

func (p *Post) GetById(ctx context.Context, id int64) (*model.PostInfo, error) {
	query := fmt.Sprintf(`SELECT * from %s where id = $1`, postgres.PostTable)

	result := p.con.QueryRowxContext(ctx, query, id)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var foundPost model.PostInfo
	err := result.StructScan(&foundPost)
	if err != nil {
		return nil, err
	}
	return &foundPost, nil
}

func (p *Post) Save(ctx context.Context, newPost *model.PostInfoInput) (int64, error) {
	query := fmt.Sprintf(`insert into %s 
			(header, text, author, time_to_live) VALUES ($1, $2, $3, $4) returning id`, postgres.PostTable)

	result := p.con.QueryRowxContext(ctx, query, newPost.Header, newPost.Text, newPost.Author, newPost.TTL)
	if result.Err() != nil {
		return 0, result.Err()
	}

	var id int64
	err := result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *Post) UpdateById(ctx context.Context, id int64, updatePost *model.PostInfoUpdate) (*model.PostInfo, error) {
	query := fmt.Sprintf(`update %s set header=$1, text=$2, time_to_live=$3 where id=$4 returning *`, postgres.PostTable)

	result := p.con.QueryRowxContext(ctx, query, updatePost.Header, updatePost.Text, updatePost.TTL, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var postAfterUpdate model.PostInfo
	err := result.StructScan(&postAfterUpdate)
	if err != nil {
		return nil, err
	}
	return &postAfterUpdate, nil
}

func (p *Post) DeleteById(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`delete from %s where id=$1`, postgres.PostTable)

	result := p.con.QueryRowxContext(ctx, query, id)
	return result.Err()
}

func (p *Post) GetAllPostTTLBefore(ctx context.Context, ttl time.Time) ([]*model.PostInfo, error) {
	query := fmt.Sprintf("Select * from %s where time_to_live < $1", postgres.PostTable)

	var postsInfo []*model.PostInfo
	err := p.con.SelectContext(ctx, &postsInfo, query, ttl)
	if err != nil {
		return nil, err
	}
	return postsInfo, nil
}
