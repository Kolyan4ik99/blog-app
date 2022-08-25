package repository

import (
	"context"
	"fmt"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/jmoiron/sqlx"
)

type PostInterface interface {
	GetAllByAuthorId(ctx context.Context, authorId int64) ([]*model.PostInfo, error)
	GetById(ctx context.Context, id int64) (*model.PostInfo, error)
	Save(ctx context.Context, newPost *model.PostInfo) (int64, error)
	UpdateById(ctx context.Context, id int64, newPost *model.PostInfo) (*model.PostInfo, error)
	DeleteById(ctx context.Context, id int64) error
}

type Post struct {
	con *sqlx.DB
}

func NewPost(con *sqlx.DB) *Post {
	return &Post{con: con}
}

func (p *Post) GetAllByAuthorId(ctx context.Context, authorId int64) ([]*model.PostInfo, error) {
	query := fmt.Sprintf(`SELECT * FROM %s where author_id = $1`, postTable)

	var postsInfo []*model.PostInfo
	err := p.con.Select(&postsInfo, query)
	return postsInfo, err
}

func (p *Post) GetById(ctx context.Context, id int64) (*model.PostInfo, error) {
	query := fmt.Sprintf(`SELECT * from %s where id = $1`, postTable)

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

func (p *Post) Save(ctx context.Context, newPost *model.PostInfo) (int64, error) {
	query := fmt.Sprintf(`insert into %s 
			(header, text, author) VALUES ($1, $2, $3) returning id`, postTable)

	result := p.con.QueryRowxContext(ctx, query, newPost.Header, newPost.Text, newPost.Author)
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

func (p *Post) UpdateById(ctx context.Context, id int64, updatePost *model.PostInfo) (*model.PostInfo, error) {
	query := fmt.Sprintf(`update %s set header=$1, text=$2 where id=$3`, postTable)

	_, err := p.con.Exec(query, updatePost.Header, updatePost.Text, id)
	return nil, err
}

func (p *Post) DeleteById(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`delete from %s where id=$1`, postTable)

	_, err := p.con.Exec(query, id)
	return err
}
