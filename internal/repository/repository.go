package repository

import (
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type Authorization interface {
	CreateUser(user *entities.User) (int, error)
	GetUser(username, password string) (*entities.User, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	UpdateUser(id int, username, password string) error
}

type Repo struct {
	Authorization
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{
		Authorization: NewAuthPostgres(db),
	}
}
