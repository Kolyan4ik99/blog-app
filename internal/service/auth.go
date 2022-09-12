package service

import (
	"context"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
)

type AuthInterface interface {
	SignUp(ctx context.Context, user *model.AuthInput) (*model.AuthOutput, error)
	SignIn(ctx context.Context, user *model.AuthInput) error
	CheckToken(ctx context.Context, token string) error
}

type Auth struct {
	repo repository.UserInterface
}

func NewAuth(repo repository.UserInterface) *Auth {
	return &Auth{repo: repo}
}

func (a *Auth) SignUp(ctx context.Context, userInput *model.AuthInput) (*model.AuthOutput, error) {
	token := util.RandomString(15)

	user := &model.UserInfo{
		Name:     userInput.Name,
		Password: userInput.Password,
		Email:    userInput.Email,
		Token:    token,
	}

	id, err := a.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &model.AuthOutput{
		Token: "Bearer " + user.Token,
		Id:    id,
	}, nil
}

func (a *Auth) SignIn(ctx context.Context, user *model.AuthInput) error {
	// TODO необходимо будет реализовать
	return nil
}

func (a *Auth) CheckToken(ctx context.Context, token string) error {
	_, err := a.repo.GetByToken(ctx, token)
	return err
}
