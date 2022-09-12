package model

import (
	"fmt"
	"time"
)

const (
	AccessYes = iota
	AccessNo
)

type Access struct {
}

type PostInfo struct {
	Id        int64      `json:"id,omitempty" db:"id"`
	Author    int64      `json:"author,omitempty" db:"author"`
	Header    string     `json:"header,omitempty" db:"header"`
	Text      string     `json:"text,omitempty" db:"text"`
	TTL       *time.Time `json:"ttl,string" db:"time_to_live" example:"2022-09-09T12:52:20.641917Z"`
	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
}

type PostInfoInput struct {
	Header string `json:"header,omitempty" db:"header" binding:"required" example:"Interesting title your post"`
	Text   string `json:"text,omitempty" db:"text" binding:"required" example:"Useful text for my post"`
	Author int64  `json:"author,omitempty" db:"author" binding:"required" example:"1"`
	TTL    string `json:"ttl,omitempty" db:"time_to_live" binding:"required" example:"2022-09-11T12:05:41+04:00"`
}

type PostInfoUpdate struct {
	Header string `json:"header,omitempty" db:"header" example:"Interesting title your post"`
	Text   string `json:"text,omitempty" db:"text" example:"Useful text for my post"`
	TTL    string `json:"ttl,omitempty" db:"time_to_live" example:"2022-10-11T12:05:41+04:00"`
}

func (p *PostInfo) String() string {
	return fmt.Sprintf("PostInfo {id=[%d], author=[%d], header=[%s], text=[%s], created_at=[%s]}",
		p.Id, p.Author, p.Header, p.Text, p.CreatedAt)
}
