package main

import (
	"context"
	"github.com/Kolyan4ik99/blog-app/internal"
	"github.com/Kolyan4ik99/blog-app/pkg/handler"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"github.com/Kolyan4ik99/blog-app/pkg/service"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	internal.InitLogger(os.Stdout)
}

func main() {
	internal.Logger.Infoln("Start application")
	ctx := context.Background()

	con, err := repository.SqlCon(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})

	if err != nil {
		internal.Logger.Fatalln("Bad connection to postgres, maybe your DB IS DOWN ", err)
	}

	defer func(con *sqlx.DB) {
		err := con.Close()
		if err != nil {
			internal.Logger.Errorln("Connection to postgres doesn't close")
		} else {
			internal.Logger.Infoln("Connection to postgres successful close")

		}
	}(con)

	h := handler.Handler{
		AuthService: &service.Auth{
			Repo: &repository.Users{Con: con},
			Ctx:  &ctx,
		},
		PostService: &service.Posts{
			Repo: &repository.Posts{Con: con},
			Ctx:  &ctx,
		},
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: h.InitRouter(),
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sign
		internal.Logger.Infoln("Application was close from System signal")
		os.Exit(1)
	}()

	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
