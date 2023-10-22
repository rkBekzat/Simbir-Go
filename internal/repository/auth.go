package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) Authorization {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := a.db.Get(&user, query, username, password)
	return user, err
}
