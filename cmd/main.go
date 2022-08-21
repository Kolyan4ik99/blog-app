package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kolyan4ik99/blog-app/internal/handler"
	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/Kolyan4ik99/blog-app/internal/repository"
	"github.com/Kolyan4ik99/blog-app/internal/service"
	"github.com/jmoiron/sqlx"
)

func init() {
	logger.InitLogger(os.Stdout)
}

func main() {
	logger.Logger.Infoln("Start application")
	ctx := context.Background()

	con, err := repository.SqlCon(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
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

	h := handler.NewHandler(
		&service.Auth{
			Repo: &repository.Users{Con: con},
			Ctx:  &ctx,
		},
		&service.Posts{
			Repo: &repository.Posts{Con: con},
			Ctx:  &ctx,
		})

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sign
		logger.Logger.Infoln("Application was close from System signal")
		os.Exit(1)
	}()

	h.InitRouter().Run(":8080")

}
