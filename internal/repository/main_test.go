package repository

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Kolyan4ik99/blog-app/config"
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
)

var ctx context.Context

var userRepository *User
var postRepository *Post

func TestMain(t *testing.M) {
	ctx = context.Background()
	cfg, err := config.Init(config.QaEnv, "config")

	configDB := postgres.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Dbname:   cfg.DB.Dbname,
	}

	con, err := postgres.SqlCon(ctx, configDB)
	if err != nil {
		log.Fatal("connection to test DB is failure", err)
	}

	userRepository = NewUser(con)
	postRepository = NewPost(con)

	os.Exit(t.Run())
}
