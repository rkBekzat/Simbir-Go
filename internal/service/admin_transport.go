package service

import "vtb_api/internal/entities"

func (a *admin) GetListOfTransports(start, count int, transportType string) ([]entities.Transport, error) {
	return a.repo.GetListOfTransports(start, count, transportType)
}

func (a *admin) GetTransportById(id int) (*entities.Transport, error) {
	return a.repo.GetTransportById(id)
}

func (a *admin) CreateTransport(tr *entities.Transport) (int, error) {
	return a.repo.CreateTransport(tr)
}

func (a *admin) UpdateTransport(tr *entities.Transport) error {
	return a.repo.UpdateTransport(tr)
}

func (a *admin) DeleteTransport(id int) error {
	return a.repo.DeleteTransport(id)
}
