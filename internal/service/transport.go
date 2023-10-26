package service

import (
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

type transport struct {
	repo repository.Transport
}

func NewTransport(repo repository.Transport) Transport {
	return &transport{repo: repo}
}

func (t *transport) GetById(id int) (*entities.Transport, error) {
	return t.repo.GetById(id)
}

func (t *transport) AddTransport(tr *entities.Transport) (int, error) {
	return t.repo.AddTransport(tr)
}

func (t *transport) Update(id int, tr *entities.Transport) error {
	return t.repo.Update(id, tr)
}

func (t *transport) Delete(id int) error {
	return t.repo.Delete(id)
}
