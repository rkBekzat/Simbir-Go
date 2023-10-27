package service

import (
	"errors"
	"fmt"
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

func (t *transport) AddTransport(ownerId int, tr *entities.Transport) (int, error) {
	return t.repo.AddTransport(ownerId, tr)
}

func (t *transport) Update(ownerId int, tr *entities.Transport) error {
	newT, err := t.repo.GetById(tr.Id)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Transport with %d id not exist, error: %s",
			tr.Id, err.Error()))
	}
	if newT.OwnerId != ownerId {
		return errors.New("User not owner of transport")
	}
	return t.repo.Update(tr)
}

func (t *transport) Delete(ownerId, id int) error {
	newT, err := t.repo.GetById(id)
	if err != nil {
		return errors.New(fmt.Sprintf(
			"Transport with %d id not exist, error: %s",
			newT.Id, err.Error()))
	}
	if newT.OwnerId != ownerId {
		return errors.New("User not owner of transport")
	}
	return t.repo.Delete(id)
}
