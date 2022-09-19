package transport

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Kolyan4ik99/blog-app/internal/model"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/Kolyan4ik99/blog-app/pkg/util"
)

var mockServer http.Handler
var testAuthorId int64
var testAuthorToken string

func TestMain(t *testing.M) {
	ctx := context.Background()

	// DI auth
	mockAuthRepository := repository.NewUser()
	mockAuthService := service.NewAuth(mockAuthRepository)
	mockAuthTransport := NewAuth(ctx, mockAuthService)

	// DI post
	mockPostRepository := repository.NewPost()
	mockPostService := service.NewPost(mockPostRepository)
	mockPostTransport := NewPost(ctx, mockPostService)

	// DI access
	mockAccessRepository := repository.NewAccess()
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
