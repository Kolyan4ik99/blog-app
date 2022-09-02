package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	UserTable = "users"
	PostTable = "posts"
)

type Config struct {
	Host,
	Port,
	User,
	Password,
	SSLMode,
	Dbname string
}

func SqlCon(ctx context.Context, c Config) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			c.Host, c.Port, c.User, c.Password, c.Dbname, c.SSLMode))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return db, nil
}
