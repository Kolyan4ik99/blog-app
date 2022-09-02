package repository

import (
	"database/sql"
	"testing"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *model.UserInfo {
	arg := &model.UserInfo{
		Name:     util.RandomString(6),
		Email:    randomEmail(),
		Password: util.RandomString(6),
	}

	id, err := TestUserRepository.Create(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, id)

	newUser, err := TestUserRepository.GetById(ctx, id)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, id, newUser.Id)
	require.Equal(t, arg.Name, newUser.Name)
	require.Equal(t, arg.Email, newUser.Email)
	require.Equal(t, arg.Password, newUser.Password)
	return newUser
}

func TestUser_Save_GetById(t *testing.T) {
	createRandomUser(t)
}

func TestUser_DeleteById(t *testing.T) {
	user := createRandomUser(t)

	err := TestUserRepository.DeleteById(ctx, user.Id)
	require.NoError(t, err)

	deleteUser, err := TestUserRepository.GetById(ctx, user.Id)
	require.Empty(t, deleteUser)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func TestUser_UpdateById(t *testing.T) {
	user := createRandomUser(t)

	argsForUpdate := &model.UserInfo{
		Name:     util.RandomString(8),
		Email:    randomEmail(),
		Password: util.RandomString(8),
	}

	updateUser, err := TestUserRepository.UpdateById(ctx, user.Id, argsForUpdate)
	require.NoError(t, err)
	require.NotEmpty(t, updateUser)

	require.Equal(t, user.Id, updateUser.Id)
	require.Equal(t, argsForUpdate.Name, updateUser.Name)
	require.Equal(t, argsForUpdate.Email, updateUser.Email)
	require.Equal(t, argsForUpdate.Password, updateUser.Password)
}

func randomEmail() string {
	return util.RandomString(6) + "@email.com"
}
