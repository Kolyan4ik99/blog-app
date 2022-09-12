package repository

import (
	"context"
	"fmt"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type AccessInterface interface {
	GetAllByPostId(ctx context.Context, postId int64) ([]*model.AccessInfo, error)
	SaveAllAccess(ctx context.Context, postId int64, accessLvl string) error
	SaveAccess(ctx context.Context, postId int64, userId int64, isAccess string) error
	DeleteAllAccess(ctx context.Context, postId int64) error
}

type Access struct {
	con *sqlx.DB
}

func NewAccess(con *sqlx.DB) *Access {
	return &Access{con: con}
}

func (a *Access) GetAllByPostId(ctx context.Context, postId int64) ([]*model.AccessInfo, error) {
	query := fmt.Sprintf("select * from %s where post_id = $1", postgres.AccessTable)

	var accessInfo []*model.AccessInfo
	err := a.con.SelectContext(ctx, &accessInfo, query, postId)
	if err != nil {
		return nil, err
	}
	return accessInfo, nil
}

func (a *Access) SaveAllAccess(ctx context.Context, postId int64, accessLvl string) error {
	query := fmt.Sprintf("insert into %s (post_id, user_id, access) values ($1, $2, $3)", postgres.AccessTable)

	result := a.con.QueryRowxContext(ctx, query, postId, 0, accessLvl)
	return result.Err()
}

func (a *Access) SaveAccess(ctx context.Context, postId int64, userId int64, isAccess string) error {
	query := fmt.Sprintf("insert into %s (post_id, user_id, access) values ($1, $2, $3)", postgres.AccessTable)

	result := a.con.QueryRowxContext(ctx, query, postId, userId, isAccess)
	return result.Err()
}

func (a *Access) DeleteAllAccess(ctx context.Context, postId int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = $1", postgres.AccessTable)

	result := a.con.QueryRowxContext(ctx, query, postId)
	return result.Err()
}
