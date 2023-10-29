package repository

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

func (a *admin) NewRent() {

}

func (a *admin) EndRent() {

}

func (a *admin) UpdateRent() {

}

func (a *admin) DeleteRent() {

}
