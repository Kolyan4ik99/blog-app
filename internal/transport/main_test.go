package transport

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Kolyan4ik99/blog-app/config"
	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
	"github.com/gin-gonic/gin"
)

var mockServer *gin.Engine
var testAuthorId int64
var testAuthorToken string

func TestMain(t *testing.M) {
	ctx := context.Background()
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

	// DI auth
	mockAuthRepository := repository.NewUser(con)
	mockAuthService := service.NewAuth(mockAuthRepository)
	mockAuthTransport := NewAuth(ctx, mockAuthService)

	// DI post
	mockPostRepository := repository.NewPost(con)
	mockPostService := service.NewPost(mockPostRepository)
	mockPostTransport := NewPost(ctx, mockPostService)

	// DI access
	mockAccessRepository := repository.NewAccess(con)
	mockAccessService := service.NewAccess(mockAccessRepository)
	mockAccessTransport := NewAccess(ctx, mockAccessService)

	mockServer = NewHandler(
		mockAuthTransport,
		mockPostTransport,
		mockAccessTransport).InitRouter()

	user, err := mockAuthService.SignUp(ctx, &model.AuthInput{
		Name:     "TestPolzak",
		Password: util.RandomString(10),
		Email:    util.RandomString(6) + "@gmail",
	})
	if err != nil {
		log.Fatal("Some problems with user", err)
	}
	testAuthorId = user.Id
	testAuthorToken = user.Token

	os.Exit(t.Run())
}
