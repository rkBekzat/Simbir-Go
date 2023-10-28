package service

import "vtb_api/internal/repository"

type admin struct {
	repo repository.Admin
}

func NewAdmin(repo repository.Admin) Admin {
	return &admin{repo: repo}
}

func (a *admin) GetAccounts() {

}

func (a *admin) GetAccountById() {

}

func (a *admin) CreateAccount() {

}

func (a *admin) UpdateAccount() {

}

func (a *admin) DeleteAccount() {

}
