package repository

import (
	"fmt"
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
	var transp entities.Transport
	query := fmt.Sprintf("SELECT id, owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price FROM %s WHERE id=$1", transportTable)
	err := t.db.Get(&transp, query, id)
	return &transp, err
}

func (t *transport) AddTransport(ownerId int, tr *entities.Transport) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s "+
		"(owner_id, can_be_rented, transport_type, model, color, identifier, description, latitude, longitude, minute_price, day_price) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)  RETURNING id", transportTable)
	row := t.db.QueryRow(query, ownerId, tr.CanBeRented, tr.TransportType, tr.Model, tr.Color, tr.Identifier, tr.Description, tr.Latitude, tr.Longitude, tr.MinutePrice, tr.DayPrice)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (t *transport) Update(tr *entities.Transport) error {
	query := fmt.Sprintf("UPDATE %s SET can_be_rented=$1, model=$2, color=$3, identifier=$4, description=$5, latitude=$6, longitude=$7, minute_price=$8, day_price=$9 WHERE id=$10", transportTable)
	_, err := t.db.Exec(query, tr.CanBeRented, tr.Model, tr.Color, tr.Identifier, tr.Description, tr.Latitude, tr.Longitude, tr.MinutePrice, tr.DayPrice, tr.Id)
	return err
}

func (t *transport) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", transportTable)
	_, err := t.db.Exec(query, id)
	return err
}
