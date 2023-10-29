package repository

import (
	"fmt"
	"time"
	"vtb_api/internal/entities"
)

func (a *admin) GetRentById(id int) (*entities.Rent, error) {
	var r entities.Rent
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", rentTable)
	err := a.db.Get(&r, query, id)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (a *admin) GetUserHistory(id int) ([]entities.Rent, error) {
	var result []entities.Rent
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", rentTable)
	err := a.db.Select(&result, query, id)
	return result, err
}

func (a *admin) GetTransportHistory(id int) ([]entities.Rent, error) {
	var result []entities.Rent
	query := fmt.Sprintf("SELECT * FROM %s WHERE transport_id=$1", rentTable)
	err := a.db.Select(&result, query, id)
	return result, err
}

func (a *admin) NewRent(r *entities.Rent) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s "+
		"(transport_id, user_id, price_of_unit, price_type, time_start) "+
		"VALUES () RETURNING id", rentTable)
	row := a.db.QueryRow(query, r.TransportId, r.UserId, r.PriceOFUnit, r.PriceType, time.Now())
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *admin) EndRent(id int, finalPrice float64) error {
	query := fmt.Sprintf("UPDATE %s SET time_end=$1, final_price=$2 WHERE id=$3", rentTable)
	_, err := a.db.Exec(query, time.Now(), finalPrice, id)
	return err
}

func (a *admin) UpdateRent(r *entities.Rent) error {
	query := fmt.Sprintf("UPDATE %s SET "+
		"transport_id=$1, user_id=$2, price_of_unit=$3, price_type=$4, time_start=$5, time_end=$6, final_price=$7 "+
		"WHERE id=$8", rentTable)
	_, err := a.db.Exec(query, r.TransportId, r.UserId, r.PriceOFUnit, r.PriceType, r.TimeStart, r.TimeEnd, r.FinalPrice, r.Id)
	return err
}

func (a *admin) DeleteRent(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", rentTable)
	_, err := a.db.Exec(query, id)
	return err
}
