package repository

import (
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

	id, err := postRepository.Save(ctx, argsPost)
	require.NoError(t, err)
	require.NotZero(t, id)

	createPost, err := postRepository.GetById(ctx, id)
	require.NoError(t, err)
	require.NotEmpty(t, createPost)

	require.Equal(t, id, createPost.Id)
	require.Equal(t, argsPost.Author, createPost.Author)
	require.Equal(t, argsPost.Header, createPost.Header)
	require.Equal(t, argsPost.Text, createPost.Text)

	return createPost
}

func TestGetById(t *testing.T) {
	user := createRandomUser(t)

	createRandomPost(t, user)
}
