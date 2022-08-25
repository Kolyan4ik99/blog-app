package repository

import (
	"context"
	"fmt"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	GetById(ctx context.Context, id int64) (*model.UserInfo, error)
	Save(ctx context.Context, newUser *model.UserInfo) (int64, error)
	DeleteById(ctx context.Context, userId int64) error
	UpdateById(ctx context.Context, userId int64, user *model.UserInfo) (*model.UserInfo, error)
}

type User struct {
	con *sqlx.DB
}

func NewUser(newCon *sqlx.DB) *User {
	return &User{
		con: newCon,
	}
}

// GetById find user in table by id
func (u *User) GetById(ctx context.Context, id int64) (*model.UserInfo, error) {
	query := fmt.Sprintf(`select * from %s where id = $1`, userTable)

	result := u.con.QueryRowxContext(ctx, query, id)
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user model.UserInfo
	err := result.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save insert user to table with users
func (u *User) Save(ctx context.Context, newUser *model.UserInfo) (int64, error) {
	query := fmt.Sprintf(`insert into %s (name, password, email) values ($1, $2, $3) returning id`, userTable)

	rows := u.con.QueryRowxContext(ctx, query, newUser.Name, newUser.Password, newUser.Email)
	if rows.Err() != nil {
		return 0, rows.Err()
	}

	var newUserId int64
	err := rows.Scan(&newUserId)
	if err != nil {
		return 0, err
	}
	return newUserId, nil
}

func (u *User) DeleteById(ctx context.Context, userId int64) error {
	query := fmt.Sprintf(`delete from %s where id=$1`, userTable)

	resultDB := u.con.QueryRowxContext(ctx, query, userId)
	if resultDB.Err() != nil {
		return resultDB.Err()
	}
	return nil
}

func (u *User) UpdateById(ctx context.Context, userId int64, user *model.UserInfo) (*model.UserInfo, error) {
	query := fmt.Sprintf(`update %s set name=$1, password=$2, email=$3 where id=$4 returning *`, userTable)

	resultDB := u.con.QueryRowxContext(ctx, query, user.Name, user.Password, user.Email, userId)
	if resultDB.Err() != nil {
		return nil, resultDB.Err()
	}
	var updateUser model.UserInfo
	err := resultDB.StructScan(&updateUser)
	if err != nil {
		return nil, err
	}
	return &updateUser, nil
}
