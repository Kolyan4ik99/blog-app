package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Kolyan4ik99/blog-app/internal"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	Con *sqlx.DB
}

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

func (u *Users) Find(ctx *context.Context, id int) *UserInfo {
	// TODO придумать по какому токену идентифицировать пользователя
	// TODO реализовать поиск пользователя по токену
	return nil
}

// Save сохраняем пользователя в таблице
func (u *Users) Save(ctx *context.Context, info *UserInfo) (*UserInfo, error) {
	countQuery := fmt.Sprintf(`select count(*) from %s where name=$1 and email=$2`, userTable)

	var counts int
	err := u.Con.Get(&counts, countQuery, info.Name, info.Email)
	if err != nil {
		return nil, err
	}

	if counts > 0 {
		return nil, errors.New(fmt.Sprintf("User with name=%s and email=%s is duplicate", info.Name, info.Email))
	}

	query := fmt.Sprintf(`insert into %s (name, password, email) values ($1, $2, $3)`, userTable)

	tx, err := u.Con.Begin()

	_, err = u.Con.Exec(query, info.Name, info.Password, info.Email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	findQuery := fmt.Sprintf(`select * from %s where name=$1 and email=$2`, userTable)
	var retUser UserInfo
	err = u.Con.Get(&retUser, findQuery, info.Name, info.Email)
	fmt.Println("AAA", retUser)
	if err != nil {
		internal.Logger.Infoln("AAA")
		return nil, err
	}
	return &retUser, nil
}
