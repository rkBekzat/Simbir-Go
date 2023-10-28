package repository

import "github.com/jmoiron/sqlx"

type admin struct {
	db *sqlx.DB
}

func NewAdmin(db *sqlx.DB) Admin {
	return &admin{db: db}
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
