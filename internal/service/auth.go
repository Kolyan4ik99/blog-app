package service

import (
	"context"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type AuthInterface interface {
	SignUp(ctx context.Context, user *model.UserInfo) (int64, error)
	SignIn(ctx context.Context, user *model.UserInfo) error
}

type Auth struct {
	repo repository.UserInterface
}

func NewAuth(repo repository.UserInterface) *Auth {
	return &Auth{repo: repo}
}

func (a *Auth) SignUp(ctx context.Context, user *model.UserInfo) (int64, error) {
	newUserId, err := a.repo.Save(ctx, user)
	if err != nil {
		return 0, err
	}

	return newUserId, nil
}

func (a *Auth) SignIn(ctx context.Context, user *model.UserInfo) error {
	// TODO необходимо реализовать аутентификацию
	panic("implement me")
	return nil
}
