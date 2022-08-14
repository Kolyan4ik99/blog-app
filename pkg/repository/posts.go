package repository

import (
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

func (p *Posts) GetAll() []*PostInfo {
	post := new(PostInfo)

	post.Author = 124
	post.Id = 1526
	post.Header = "Название статьи"
	post.Text = "Тело статьи. ААА"

	ret := make([]*PostInfo, 1)
	ret[0] = post
	return ret
}

func (p *Posts) GetById(id int) (*PostInfo, error) {
	query := fmt.Sprintf(`SELECT * from %s where id = $1`, postTable)
	var postInst PostInfo
	err := p.Con.Get(&postInst, query, id)
	if err != nil {
		return nil, err
	}

	return &postInst, nil
}

func (p *Posts) Save() {
}

func (p *Posts) UpdateById() {
}

func (p *Posts) DeleteById() {
}
