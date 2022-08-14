package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Users struct {
	Con *sqlx.DB
}

type UserInfo struct {
	Id       int64  `json:"-" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func (u *Users) Find(id int) {

}

func (u *Users) Save(info *UserInfo) error {
	findQuery := fmt.Sprintf(`select count(*) from %s where name=$1 and email=$2`, userTable)

	var counts int
	err := u.Con.Get(&counts, findQuery, info.Name, info.Email)
	if err != nil {
		return err
	}

	if counts > 0 {
		return errors.New(fmt.Sprintf("User with name=%s and email=%s is duplicate", info.Name, info.Email))
	}

	query := fmt.Sprintf(`insert into %s (name, password, email) values ($1, $2, $3)`, userTable)

	tx, err := u.Con.Begin()

	result, err := u.Con.Exec(query, info.Name, info.Password, info.Email)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(result)
	tx.Commit()
	return nil
}
