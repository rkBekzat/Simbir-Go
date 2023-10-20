package service

import "vtb_api/internal/repository"

type Auth interface {
}

type UseCase struct {
	Auth
}

func NewUseCase(repo repository.Repo) *UseCase {
	return &UseCase{}
}
