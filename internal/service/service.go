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

type UseCase struct {
	Auth
	Transport
	Renting
}

func NewUseCase(repo repository.Repo) *UseCase {
	return &UseCase{
		Auth:      NewAuthorization(repo.Authorization),
		Transport: NewTransport(repo.Transport),
		Renting:   NewRent(repo.Renting, repo.Transport),
	}
}
