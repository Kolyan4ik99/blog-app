package repository

import (
	"context"
	"errors"

	"github.com/Kolyan4ik99/blog-app/internal/model"
)

type AccessInterface interface {
	GetAllByPostId(ctx context.Context, postId int64) ([]*model.AccessInfo, error)
	SaveAllAccess(ctx context.Context, postId int64, accessLvl string) error
	SaveAccess(ctx context.Context, postId int64, userId int64, isAccess string) error
	DeleteAllAccess(ctx context.Context, postId int64) error
}

type Access struct {
	info map[int64][]*model.AccessInfo
}

type AccessesInfo struct {
	info []*model.AccessInfo
}

func NewAccess() *Access {
	return &Access{
		info: make(map[int64][]*model.AccessInfo, 0),
	}
}

func (a *Access) GetAllByPostId(ctx context.Context, postId int64) ([]*model.AccessInfo, error) {
	infos, exist := a.info[postId]
	if !exist {
		return nil, errors.New("bad post_id")
	}
	return infos, nil
}

func (a *Access) SaveAllAccess(ctx context.Context, postId int64, accessLvl string) error {
	a.info[postId] = make([]*model.AccessInfo, 0)
	a.info[postId] = append(a.info[postId], &model.AccessInfo{PostId: postId, Access: accessLvl, UserId: 0})
	return nil
}

func (a *Access) SaveAccess(ctx context.Context, postId int64, userId int64, isAccess string) error {
	_, exist := a.info[postId]
	if !exist {
		a.info[postId] = make([]*model.AccessInfo, 0)
	}
	a.info[postId] = append(a.info[postId], &model.AccessInfo{PostId: postId, Access: isAccess, UserId: userId})
	return nil
}

func (a *Access) DeleteAllAccess(ctx context.Context, postId int64) error {
	delete(a.info, postId)
	return nil
}
