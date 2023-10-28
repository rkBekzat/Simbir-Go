package service

import (
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

type admin struct {
	repo repository.Admin
}

func NewAdmin(repo repository.Admin) Admin {
	return &admin{repo: repo}
}

func (a *admin) GetAccounts(start, count int) ([]entities.User, error) {
	return a.repo.GetAccounts(start, count)
}

func (a *admin) GetAccountById(id int) (*entities.User, error) {
	return a.repo.GetAccountById(id)
}

func (a *admin) CreateAccount(user *entities.User) (int, error) {
	return a.repo.CreateAccount(user)
}

func (a *admin) UpdateAccount(user *entities.User) error {
	return a.repo.UpdateAccount(user)
}

func (a *admin) DeleteAccount(id int) error {
	return a.repo.DeleteAccount(id)
}
