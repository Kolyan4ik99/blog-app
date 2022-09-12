package model

type AccessInfo struct {
	PostId int64  `json:"post_id,omitempty" db:"post_id" example:"1"`
	UserId int64  `json:"user_id,omitempty" db:"user_id" example:"1"`
	Access string `json:"access,omitempty" db:"access" example:"yes, no"`
}

type AccessSetInput struct {
	UsersId []IdAndType
	IsAll   string `json:"is_all,omitempty" example:"yes"`
}

type AccessOutput struct {
	PostId  int64 `json:"post_id,omitempty" example:"1"`
	UsersId []IdAndType
	IsAll   string `json:"is_all,omitempty" example:"yes"`
}

type IdAndType struct {
	Id   int64  `json:"id,omitempty" binding:"required" example:"1"`
	Type string `json:"type,omitempty" binding:"required" example:"yes"`
}
