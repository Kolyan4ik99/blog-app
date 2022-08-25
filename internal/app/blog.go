package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kolyan4ik99/blog-app/config"
	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/Kolyan4ik99/blog-app/internal/transport"
	"github.com/Kolyan4ik99/blog-app/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

func Run() {
	logger.Logger.Infoln("Start application")
	cfg, err := config.Init(config.ProdEnv, "config")
	if err != nil {
		log.Fatal("bad initial config file", err)
	}
	ctx := context.Background()

	con, err := postgres.SqlCon(ctx, postgres.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Dbname:   cfg.DB.Dbname,
		SSLMode:  cfg.DB.SSLMode,
	})

	if err != nil {
		logger.Logger.Fatalln("Bad connection to postgres, maybe your DB IS DOWN ", err)
	}

	defer func(con *sqlx.DB) {
		err := con.Close()
		if err != nil {
			logger.Logger.Errorln("Connection to postgres doesn't close")
		} else {
			logger.Logger.Infoln("Connection to postgres successful close")

		}
	}(con)

	// DI auth
	authRepository := repository.NewUser(con)
	authService := service.NewAuth(authRepository)
	authTransport := transport.NewAuth(ctx, authService)

	// DI post
	postRepository := repository.NewPost(con)
	postService := service.NewPost(postRepository)
	postTransport := transport.NewPost(ctx, postService)

	h := transport.NewHandler(
		authTransport,
		postTransport)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sign
		logger.Logger.Infoln("Application was close from System signal")
		os.Exit(1)
	}()

	err = h.InitRouter().Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return
	}
}
