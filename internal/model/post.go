package model

import (
	"fmt"
	"time"
)

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
