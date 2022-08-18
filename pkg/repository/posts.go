package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Posts struct {
	Con *sqlx.DB
}

type PostInfo struct {
	Id        int64     `json:"-" db:"id"`
	Author    int64     `json:"author,omitempty" db:"author"`
	Header    string    `json:"header,omitempty" db:"header"`
	Text      string    `json:"text,omitempty" db:"text"`
	CreatedAt time.Time `json:"-" db:"created_at"`
}

func (p *PostInfo) String() string {
	return fmt.Sprintf("PostInfo {id=[%d], author=[%d], header=[%s], text=[%s], created_at=[%s]}",
		p.Id, p.Author, p.Header, p.Text, p.CreatedAt)
}

func (p *Posts) GetAll(ctx *context.Context) ([]*PostInfo, error) {
	query := fmt.Sprintf(`SELECT * FROM %s`, postTable)

	var postsInfo []*PostInfo
	err := p.Con.Select(&postsInfo, query)
	return postsInfo, err
}

func (p *Posts) GetById(ctx *context.Context, id int) (*PostInfo, error) {
	query := fmt.Sprintf(`SELECT * from %s where id = $1`, postTable)

	var postInst PostInfo
	err := p.Con.Get(&postInst, query, id)
	if err != nil {
		return nil, err
	}

	return &postInst, nil
}

func (p *Posts) Save(ctx *context.Context, newPost *PostInfo) error {
	query := fmt.Sprintf(`insert into %s 
			(header, text, author) VALUES ($1, $2, $3)`, postTable)

	_, err := p.Con.Exec(query, newPost.Header, newPost.Text, newPost.Author)
	return err
}

func (p *Posts) UpdateById(ctx *context.Context, newPost *PostInfo, id int) error {
	query := fmt.Sprintf(`update %s set header=$1, text=$2 where id=$3`, postTable)

	_, err := p.Con.Exec(query, newPost.Header, newPost.Text, id)
	return err
}

func (p *Posts) DeleteById(ctx *context.Context, id int) error {
	query := fmt.Sprintf(`delete from %s where id=$1`, postTable)

	_, err := p.Con.Exec(query, id)
	return err
}
