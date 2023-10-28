package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type rent struct {
	db *sqlx.DB
}

func NewRent(db *sqlx.DB) Renting {
	return &rent{db: db}
}

func (r *rent) AccessTransport(lat, long, radius float64, tp string) ([]int, error) {
	var result []int
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE "+
			"(latitude BETWEEN $1 AND $2) AND "+
			"(longitude BETWEEN $3 AND $4) AND "+
			"transport_type=$5",
		transportTable)
	err := r.db.Select(&result, query, lat-radius, lat+radius, long-radius, long+radius, tp)
	return result, err
}

func (r *rent) GetById(id int) (*entities.Rent, error) {
	var result entities.Rent
	query := fmt.Sprintf("SELECT id, transport_id, user_id FROM %s WHERE id=$1", rentTable)
	err := r.db.Get(&result, query, id)
	return &result, err
}

func (r *rent) History(userId int) ([]entities.Rent, error) {
	var result []entities.Rent
	query := fmt.Sprintf("SELECT id, transport_id, user_id FROM %s WHERE user_id=$1", rentTable)
	err := r.db.Select(&result, query, userId)
	return result, err
}

func (r *rent) TransportHistory(transportId int) ([]entities.Rent, error) {
	var result []entities.Rent
	query := fmt.Sprintf("SELECT id, transport_id, user_id FROM %s WHERE transport_id=$1", rentTable)
	err := r.db.Select(&result, query, transportId)
	return result, err
}

func (r *rent) StartRenting(userId, transportID int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, tranposrt_id) VALUES ($1, $2) RETURNING id", rentTable)
	row := r.db.QueryRow(query, userId, transportID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	query = fmt.Sprintf("UPDATE %s SET can_be_rentend=false", transportTable)
	_, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *rent) EndRenting(transportId int, lat, long float64) error {
	query := fmt.Sprintf("UPDATE %s SET can_be_rentend=true, latitude=$2, longtitude=$3 WHERE id=$1", transportTable)
	_, err := r.db.Exec(query, transportId, lat, long)
	return err
}
