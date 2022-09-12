package service

import (
	"context"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
)

type AccessInterface interface {
	GetAccessPost(ctx context.Context, postId int64) (*model.AccessOutput, error)
	SetAccessPost(ctx context.Context, postId int64, input *model.AccessSetInput) error
}

type Access struct {
	repo repository.AccessInterface
}

func NewAccess(repo repository.AccessInterface) *Access {
	return &Access{repo: repo}
}

func (a *Access) GetAccessPost(ctx context.Context, postId int64) (*model.AccessOutput, error) {
	accessPosts, err := a.repo.GetAllByPostId(ctx, postId)
	if err != nil {
		return nil, err
	}
	retAccess := &model.AccessOutput{
		PostId:  postId,
		UsersId: make([]model.IdAndType, len(accessPosts)),
	}
	for i, info := range accessPosts {
		retAccess.UsersId[i].Id = info.UserId
		retAccess.UsersId[i].Type = info.Access
	}
	if len(accessPosts) == 0 {
		retAccess.IsAll = "no"
	}
	if len(accessPosts) == 1 && accessPosts[0].UserId == 0 {
		retAccess.IsAll = accessPosts[0].Access
	}
	return retAccess, nil
}

// SetAccessPost если изменения прав у всех пользователей поста, то мы можем удалить старые права.
// После удаления и изменения, в запросе  могут быть пользаки с противоположным состоянием. Добавляем их так-же
// Если изменения не касаются всех, то добавляются новые записи (могут возникнуть дубликаты, починить потом)
func (a *Access) SetAccessPost(ctx context.Context, postId int64, input *model.AccessSetInput) error {
	if input.IsAll != "" {
		accessType := input.IsAll

		err := a.repo.DeleteAllAccess(ctx, postId)
		if err != nil {
			return err
		}

		err = a.repo.SaveAllAccess(ctx, postId, accessType)
		if err != nil {
			return err
		}
		for _, info := range input.UsersId {
			if info.Type != accessType {
				err = a.repo.SaveAccess(ctx, postId, info.Id, info.Type)
				if err != nil {
					return err
				}
			}
		}
	} else {
		// TODO fix дубликатов быть не может, из-за constraint в DB,
		for _, info := range input.UsersId {
			err := a.repo.SaveAccess(ctx, postId, info.Id, info.Type)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
