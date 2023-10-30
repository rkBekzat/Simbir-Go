package service

import "vtb_api/internal/entities"

func (a *admin) GetRentById(id int) (*entities.Rent, error) {
	return nil, nil
}

func (a *admin) GetUserHistory(id int) ([]entities.Rent, error) {
	return nil, nil
}

func (a *admin) GetTransportHistory(id int) ([]entities.Rent, error) {
	return nil, nil
}

func (a *admin) NewRent(r *entities.Rent) (int, error) {
	return a.repo.NewRent(r)
}

func (a *admin) EndRent(id int, lat, long float64) error {
	var finalPrice float64
	return a.repo.EndRent(id, finalPrice)
}

func (a *admin) UpdateRent(r *entities.Rent) error {
	return a.repo.UpdateRent(r)
}

func (a *admin) DeleteRent(id int) error {
	return a.repo.DeleteRent(id)
}
