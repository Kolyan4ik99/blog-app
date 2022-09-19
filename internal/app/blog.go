package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Kolyan4ik99/blog-app/docs"
	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/Kolyan4ik99/blog-app/internal/transport"
)

// @title           Blog API
// @version         1.0
// @description     REST API for Blog App.
// @host      localhost:8080
// @BasePath  /v1/

// @securityDefinitions.apikey ApiKeyAuth
// @name                       Authorization
// @in                         header
// @description                Example: Bearer token

func Run() {
	logger.Logger.Infoln("Start application")
	//cfg, err := config.Init(config.ProdEnv, "config")
	//if err != nil {
	//	log.Fatal("bad initial config file", err)
	//}
	ctx := context.Background()

	// DI auth
	authRepository := repository.NewUser()
	authService := service.NewAuth(authRepository)
	authTransport := transport.NewAuth(ctx, authService)

	// DI post
	postRepository := repository.NewPost()
	postService := service.NewPost(postRepository)
	postTransport := transport.NewPost(ctx, postService)

	accessRepository := repository.NewAccess()
	accessService := service.NewAccess(accessRepository)
	accessTransport := transport.NewAccess(ctx, accessService)

	h := transport.NewHandler(
		authTransport,
		postTransport,
		accessTransport)

	scheduler := NewTimeExpiryScheduler(postRepository, time.Minute)
	go scheduler.Run(ctx) // Start async

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sign
		logger.Logger.Infoln("Application was close from System signal")
		os.Exit(1)
	}()

	//serv := http.Server{
	//	Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
	//	Handler:        h.InitRouter(),
	//	ReadTimeout:    60 * time.Second,
	//	WriteTimeout:   60 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	http.ListenAndServe(":8080", h.InitRouter())

	//err = serv.ListenAndServe()
	//if err != nil {
	//	return
	//}
}
