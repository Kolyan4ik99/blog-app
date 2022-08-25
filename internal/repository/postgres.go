package repository

import (
	"context"
	"fmt"

	"github.com/Kolyan4ik99/blog-app/internal/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable = "users"
	postTable = "posts"
)

type Config struct {
	Host,
	Port,
	User,
	Password,
	Dbname string
}

func SqlCon(ctx context.Context, c Config) (*sqlx.DB, error) {
	logger.Logger.Infoln("Try connect to DataBase")
	db, err := sqlx.ConnectContext(ctx, "postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.Dbname))
	if err != nil {
		return nil, err
	}
	logger.Logger.Infoln("Connection is successful!")

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return db, nil
}
