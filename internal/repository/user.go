package repository

import (
	"context"
	"database/sql"

	"github.com/Kolyan4ik99/blog-app/internal/model"
)

type UserInterface interface {
	GetById(ctx context.Context, id int64) (*model.UserInfo, error)
	Create(ctx context.Context, newUser *model.UserInfo) (int64, error)
	DeleteById(ctx context.Context, userId int64) error
	UpdateById(ctx context.Context, userId int64, user *model.UserInfo) (*model.UserInfo, error)
	GetByToken(ctx context.Context, token string) (*model.UserInfo, error)
}

type User struct {
	maxId     int64
	usersInfo map[int64]*model.UserInfo
}

func NewUser() *User {
	return &User{
		maxId:     1,
		usersInfo: make(map[int64]*model.UserInfo, 50),
	}
}

// GetById find user in table by id
func (u *User) GetById(ctx context.Context, id int64) (*model.UserInfo, error) {
	user, exist := u.usersInfo[id]
	if !exist {
		return nil, sql.ErrNoRows
	}
	return user, nil
}

// Create insert user to table with users
func (u *User) Create(ctx context.Context, newUser *model.UserInfo) (int64, error) {
	u.usersInfo[u.maxId] = &model.UserInfo{
		Id:       u.maxId,
		Name:     newUser.Name,
		Password: newUser.Password,
		Email:    newUser.Email,
		Token:    newUser.Token,
	}
	u.maxId++
	return u.maxId - 1, nil
}

func (u *User) DeleteById(ctx context.Context, userId int64) error {
	_, exist := u.usersInfo[userId]
	if !exist {
		return sql.ErrNoRows
	}
	delete(u.usersInfo, userId)
	return nil
}

func (u *User) UpdateById(ctx context.Context, userId int64, updateUser *model.UserInfo) (*model.UserInfo, error) {
	user, exist := u.usersInfo[userId]
	if !exist {
		return nil, sql.ErrNoRows
	}
	user.Name = updateUser.Name
	user.Password = updateUser.Password
	user.Email = updateUser.Email
	return user, nil
}

func (u *User) GetByToken(ctx context.Context, token string) (*model.UserInfo, error) {
	for _, info := range u.usersInfo {
		if info.Token == token {
			return info, nil
		}
	}
	return nil, sql.ErrNoRows
}
