package repository

import (
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
}

type Repo struct {
	Authorization
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{
		Authorization: NewAuthPostgres(db),
	}
}
