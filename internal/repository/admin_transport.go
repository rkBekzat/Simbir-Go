package repository

import (
	"fmt"
	"vtb_api/internal/entities"
)

func (a *admin) GetListOfTransports(start, count int, transportType string) ([]entities.Transport, error) {
	var result []entities.Transport
	query := fmt.Sprintf("SELECT TOP $1 * FROM %s WHERE "+
		"id >= $2 AND transport_type=$3", transportTable)
	err := a.db.Select(&result, query, count, start, transportType)
	return result, err
}

func (a *admin) GetTransportById(id int) (*entities.Transport, error) {
	var result entities.Transport
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", transportTable)
	err := a.db.Get(&result, query, id)
	return &result, err
}

func (a *admin) CreateTransport(tr *entities.Transport) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s "+
		"(owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)  RETURNING id", transportTable)
	row := a.db.QueryRow(query, tr.OwnerId, tr.CanBeRented, tr.TransportType, tr.Model, tr.Color, tr.Identifier, tr.Description, tr.Latitude, tr.Longitude, tr.MinutePrice, tr.DayPrice)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *admin) UpdateTransport(tr *entities.Transport) error {
	query := fmt.Sprintf("UPDATE %s SET can_be_rented=$1, model=$2, color=$3, identifier=$4, description=$5, latitude=$6, longitude=$7, minute_price=$8, day_price=$9, owner_id=$10 WHERE id=$11", transportTable)
	_, err := a.db.Exec(query, tr.CanBeRented, tr.Model, tr.Color, tr.Identifier, tr.Description, tr.Latitude, tr.Longitude, tr.MinutePrice, tr.DayPrice, tr.OwnerId, tr.Id)
	return err
}

func (a *admin) DeleteTransport(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", transportTable)
	_, err := a.db.Exec(query, id)
	return err
}
