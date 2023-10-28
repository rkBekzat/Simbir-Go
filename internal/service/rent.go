package service

import (
	"errors"
	"vtb_api/internal/entities"
	"vtb_api/internal/repository"
)

type rent struct {
	repo          repository.Renting
	repoTransport repository.Transport
}

func NewRent(repo repository.Renting, repo2 repository.Transport) Renting {
	return &rent{repo: repo, repoTransport: repo2}
}

func (r *rent) AccessTransport(lat, long, radius float64, tp string) ([]int, error) {
	return r.repo.AccessTransport(lat, long, radius, tp)
}

func (r *rent) GetById(userId, id int) (*entities.Rent, error) {
	renting, err := r.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if renting.UserId == userId {
		return renting, nil
	}
	tran, err := r.repoTransport.GetById(renting.TransportId)
	if err != nil {
		return nil, err
	}
	if tran.OwnerId != userId {
		return nil, errors.New("User can't access to this rent history")
	}
	return renting, nil
}

func (r *rent) History(id int) ([]entities.Rent, error) {
	return r.repo.History(id)
}

func (r *rent) TransportHistory(userId, transportId int) ([]entities.Rent, error) {
	tran, err := r.repoTransport.GetById(transportId)
	if err != nil {
		return []entities.Rent{}, err
	}
	if tran.OwnerId != userId {
		return nil, errors.New("User can't access to this rent history")
	}
	return r.repo.TransportHistory(transportId)
}

func (r *rent) StartRenting(userId, transportID int) (int, error) {
	return r.repo.StartRenting(userId, transportID)
}

func (r *rent) EndRenting(userId, rentId int) error {
	renting, err := r.repo.GetById(rentId)
	if err != nil {
		return err
	}
	if renting.UserId != userId {
		return errors.New("User Can't end the rent")
	}
	return r.repo.EndRenting(renting.TransportId)
}
