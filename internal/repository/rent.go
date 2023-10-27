package repository

import "github.com/jmoiron/sqlx"

type rent struct {
	db *sqlx.DB
}

func NewRent(db *sqlx.DB) Renting {
	return &rent{db: db}
}

func (r *rent) AccessTransport() {

}

func (r *rent) GetById(id int) {

}

func (r *rent) History() {

}

func (r *rent) TransportHistory() {

}

func (r *rent) StartRenting() {

}

func (r *rent) EndRenting() {

}
