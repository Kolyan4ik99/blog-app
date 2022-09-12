package model

import "fmt"

type UserInfo struct {
	Id       int64  `json:"-" db:"id" example:"1"`
	Name     string `json:"name,omitempty" db:"name" example:"Niko"`
	Password string `json:"password" db:"password" example:"(9)SomePwd@41"`
	Email    string `json:"email" db:"email" example:"Emal@gmail.com"`
	Token    string `json:"-" db:"token" example:"g3nkjhjonnqoknqegq="`
}

type AuthInput struct {
	Name     string `json:"name,omitempty" db:"name" binding:"required" example:"Niko"`
	Password string `json:"password" db:"password" binding:"required" example:"(9)SomePwd@41"`
	Email    string `json:"email" db:"email" binding:"required" example:"Emal@gmail.com"`
}

type AuthOutput struct {
	Id    int64  `json:"user_id,omitempty" example:"1"`
	Token string `json:"token,omitempty" example:"Bearer hherh34h3h14jhk5jhkmnkln"`
}

func (u *UserInfo) String() string {
	return fmt.Sprintf("UserInfo {id=[%d], name=[%s], email=[%s]}",
		u.Id, u.Name, u.Email)
}
