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
	AccessTransport(lat, long, radius float64, tp string) ([]int, error)
	GetById(id int) (*entities.Rent, error)
	History(id int) ([]entities.Rent, error)
	TransportHistory(transportId int) ([]entities.Rent, error)
	StartRenting(userId, transportID int, rentType string) (int, error)
	EndRenting(transportId, rentId int, lat, long float64) error
}

type Admin interface {
	GetAccounts(start, count int) ([]entities.User, error)
	GetAccountById(id int) (*entities.User, error)
	CreateAccount(user *entities.User) (int, error)
	UpdateAccount(user *entities.User) error
	DeleteAccount(id int) error

	GetListOfTransports(start, count int, transportType string) ([]entities.Transport, error)
	GetTransportById(id int) (*entities.Transport, error)
	CreateTransport(tran *entities.Transport) (int, error)
	UpdateTransport(tran *entities.Transport) error
	DeleteTransport(id int) error

	GetRentById(id int) (*entities.Rent, error)
	GetUserHistory(id int) ([]entities.Rent, error)
	GetTransportHistory(id int) ([]entities.Rent, error)
	NewRent(r *entities.Rent) (int, error)
	EndRent(id int, finalPrice float64) error
	UpdateRent(r *entities.Rent) error
	DeleteRent(id int) error
}

type Repo struct {
	Authorization
	Transport
	Renting
	Admin
}

func NewRepo(db *sqlx.DB) Repo {
	return Repo{
		Authorization: NewAuthPostgres(db),
		Transport:     NewTransport(db),
		Renting:       NewRent(db),
		Admin:         NewAdmin(db),
	}
}
