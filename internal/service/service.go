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

type UseCase struct {
	Auth
}

func NewUseCase(repo repository.Repo) *UseCase {
	return &UseCase{
		Auth: NewAuthorization(repo.Authorization),
	}
}
