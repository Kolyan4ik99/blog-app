package model

import "fmt"

type UserInfo struct {
	// Идентификатор
	Id int64 `json:"-" db:"id"`
	// Имя
	Name string `json:"name,omitempty" db:"name"`
	// Пароль
	Password string `json:"password" db:"password"`
	// Почта
	Email string `json:"email" db:"email"`
}

func (u *UserInfo) String() string {
	return fmt.Sprintf("UserInfo {id=[%d], name=[%s], email=[%s]}",
		u.Id, u.Name, u.Email)
}
