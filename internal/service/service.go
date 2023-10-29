package service

import (
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

type Auth interface {
	CreateUser(user *entities.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	Update(id int, username, password string) error
	Information(id int) (*entities.User, error)
}

type Transport interface {
	GetById(id int) (*entities.Transport, error)
	AddTransport(ownerId int, t *entities.Transport) (int, error)
	Update(ownerId int, t *entities.Transport) error
	Delete(ownerId, id int) error
}

type Renting interface {
	AccessTransport(lat, long, radius float64, tp string) ([]int, error)
	GetById(userId, id int) (*entities.Rent, error)
	History(id int) ([]entities.Rent, error)
	TransportHistory(userId, transportId int) ([]entities.Rent, error)
	StartRenting(userId, transportID int, rentingType string) (int, error)
	EndRenting(userId, transportId int, lat, long float64) error
}

type Admin interface {
	GetAccounts(start, count int) ([]entities.User, error)
	GetAccountById(id int) (*entities.User, error)
	CreateAccount(user *entities.User) (int, error)
	UpdateAccount(user *entities.User) error
	DeleteAccount(id int) error

	GetListOfTransports(start, count int, transportType string) ([]entities.Transport, error)
	GetTransportById(id int) (*entities.Transport, error)
	CreateTransport(tr *entities.Transport) (int, error)
	UpdateTransport(tr *entities.Transport) error
	DeleteTransport(id int) error

	GetRentById(id int) (*entities.Rent, error)
	GetUserHistory(id int) ([]entities.Rent, error)
	GetTransportHistory(id int) ([]entities.Rent, error)
	NewRent()
	EndRent()
	UpdateRent()
	DeleteRent()
}

type UseCase struct {
	Auth
	Transport
	Renting
	Admin
}

func NewUseCase(repo repository.Repo) *UseCase {
	return &UseCase{
		Auth:      NewAuthorization(repo.Authorization),
		Transport: NewTransport(repo.Transport),
		Renting:   NewRent(repo.Renting, repo.Transport),
		Admin:     NewAdmin(repo.Admin),
	}
}
