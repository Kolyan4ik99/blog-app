package repository

import (
	"context"
	"os"
	"testing"
)

var ctx context.Context

var TestUserRepository *User
var TestPostRepository *Post

func TestMain(t *testing.M) {
	ctx = context.Background()

	TestUserRepository = NewUser()
	TestPostRepository = NewPost()

	os.Exit(t.Run())
}
