package repository

import (
	"database/sql"
	"math/rand"
	"testing"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, user *model.UserInfo) *model.PostInfo {
	argsPost := &model.PostInfo{
		Author: user.Id,
		Header: util.RandomString(10),
		Text:   util.RandomString(50),
	}

	id, err := TestPostRepository.Save(ctx, argsPost)
	require.NoError(t, err)
	require.NotZero(t, id)

	createPost, err := TestPostRepository.GetById(ctx, id)
	require.NoError(t, err)
	require.NotEmpty(t, createPost)

	require.Equal(t, id, createPost.Id)
	require.Equal(t, argsPost.Author, createPost.Author)
	require.Equal(t, argsPost.Header, createPost.Header)
	require.Equal(t, argsPost.Text, createPost.Text)
	require.NotZero(t, createPost.CreatedAt)

	return createPost
}

func TestPost_Save_GetById(t *testing.T) {
	user := createRandomUser(t)

	createRandomPost(t, user)
}

func TestPost_GetAllByAuthorId(t *testing.T) {
	user := createRandomUser(t)

	n := 1 + rand.Intn(5)
	for i := 0; i < n; i++ {
		createRandomPost(t, user)
	}

	createPosts, err := TestPostRepository.GetAllByAuthorId(ctx, user.Id)
	require.NoError(t, err)
	require.NotEmpty(t, createPosts)

	require.Equal(t, n, len(createPosts))
}

func TestPost_DeleteById(t *testing.T) {
	user := createRandomUser(t)

	post := createRandomPost(t, user)

	err := TestPostRepository.DeleteById(ctx, post.Id)
	require.NoError(t, err)

	postAfterDelete, err := TestPostRepository.GetById(ctx, post.Id)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, postAfterDelete)
}

func TestPost_UpdateById(t *testing.T) {
	user := createRandomUser(t)

	post := createRandomPost(t, user)

	argsUpdate := &model.PostInfo{
		Header: util.RandomString(7),
		Text:   util.RandomString(11),
	}

	updatePost, err := TestPostRepository.UpdateById(ctx, post.Id, argsUpdate)
	require.NoError(t, err)
	require.NotEmpty(t, updatePost)

	require.Equal(t, post.Author, updatePost.Author)
	require.Equal(t, argsUpdate.Header, updatePost.Header)
	require.Equal(t, argsUpdate.Text, updatePost.Text)
}
