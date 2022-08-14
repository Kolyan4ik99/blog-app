package main

import (
	"github.com/Kolyan4ik99/blog-app/pkg/handler"
	"github.com/Kolyan4ik99/blog-app/pkg/repository"
	"github.com/Kolyan4ik99/blog-app/pkg/service"
	"log"
)

func main() {

	con, err := repository.SqlCon(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "postgres",
	})
	if err != nil {
		log.Fatal("bad connection to postgres", err)
	}
	defer con.Close()

	h := handler.Handler{
		AuthService: &service.Auth{
			Repo: repository.Users{Con: con},
		},

		PostService: &service.Posts{
			Repo: repository.Posts{Con: con},
		},
	}

	if err = h.InitRouter().Run(":80"); err != nil {
		return
	}

}
