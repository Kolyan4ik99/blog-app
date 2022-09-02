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

var TestUserRepository *User
var TestPostRepository *Post

func TestMain(t *testing.M) {
	ctx = context.Background()
	cfg, err := config.Init("../../"+config.QaEnv, "config")
	if err != nil {
		log.Fatal(err)
	}

	configDB := postgres.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Dbname:   cfg.DB.Dbname,
		SSLMode:  cfg.DB.SSLMode,
	}

	con, err := postgres.SqlCon(ctx, configDB)
	if err != nil {
		log.Fatal("connection to test DB is failure", err)
	}

	TestUserRepository = NewUser(con)
	TestPostRepository = NewPost(con)

	os.Exit(t.Run())
}
