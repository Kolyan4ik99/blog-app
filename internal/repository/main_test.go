package repository

import (
	"context"
	"log"
	"os"
	"testing"
)

var ctx context.Context

var userRepository *User
var postRepository *Post

func TestMain(t *testing.M) {
	ctx = context.Background()

	config := Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	}

	con, err := SqlCon(ctx, config)
	if err != nil {
		log.Fatal("connection to test DB is failure", err)
	}

	userRepository = NewUser(con)
	postRepository = NewPost(con)

	os.Exit(t.Run())
}
