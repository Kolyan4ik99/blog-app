package repository

import (
	"context"
	"errors"
	"fmt"
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
func (u *Users) Save(ctx *context.Context, info *UserInfo) (int64, error) {
	countQuery := fmt.Sprintf(`select count(*) from %s where name=$1 and email=$2`, userTable)

	var counts int
	err := u.Con.Get(&counts, countQuery, info.Name, info.Email)
	if err != nil {
		return -1, err
	}

	if counts > 0 {
		return -1, errors.New(fmt.Sprintf("User with name=%s and email=%s is duplicate", info.Name, info.Email))
	}

	query := fmt.Sprintf(`insert into %s (name, password, email) values ($1, $2, $3) returning id`, userTable)

	tx, _ := u.Con.Begin()

	result := u.Con.QueryRowContext(*ctx, query, info.Name, info.Password, info.Email)

	if result.Err() != nil {
		tx.Rollback()
		return -1, result.Err()
	}
	var newUserId int64
	result.Scan(&newUserId)
	tx.Commit()

	return newUserId, nil
}
