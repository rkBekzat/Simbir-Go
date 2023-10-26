package repository

import (
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type transport struct {
	db *sqlx.DB
}

func NewTransport(dp *sqlx.DB) Transport {
	return &transport{db: dp}
}

func (t *transport) GetById(id int) (*entities.Transport, error) {
	return nil, nil
}

func (t *transport) AddTransport(tr *entities.Transport) (int, error) {
	return 0, nil
}

func (t *transport) Update(id int, tr *entities.Transport) error {
	return nil
}

func (t *transport) Delete(id int) error {
	return nil
}
