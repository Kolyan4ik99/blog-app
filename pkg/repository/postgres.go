package repository

import (
	"fmt"
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

func SqlCon(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Password, c.Dbname))
	if err != nil {
		return nil, err
	}
	return db, nil
}
