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

type Transport interface {
	GetById(id int) (*entities.Transport, error)
	AddTransport(ownerId int, t *entities.Transport) (int, error)
	Update(t *entities.Transport) error
	Delete(id int) error
}

type Renting interface {
	AccessTransport()
	GetById(id int)
	History()
	TransportHistory()
	StartRenting()
	EndRenting()
}

type Repo struct {
	Authorization
	Transport
	Renting
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{
		Authorization: NewAuthPostgres(db),
		Transport:     NewTransport(db),
		Renting:       NewRent(db),
	}
}
