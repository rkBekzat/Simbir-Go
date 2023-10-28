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

func (a *AuthPostgres) CreateUser(user *entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, is_admin, balance) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password, user.IsAdmin, user.Balance)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (*entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := a.db.Get(&user, query, username, password)
	return &user, err
}

func (a *AuthPostgres) GetUserById(id int) (*entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT username, name, is_admin, balace FROM %s WHERE id=$1", userTable)
	row := a.db.QueryRow(query, id)
	if err := row.Scan(&user.Username, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthPostgres) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", userTable)
	err := a.db.Get(&user, query, username)
	return &user, err
}

func (a *AuthPostgres) UpdateUser(id int, username, password string) error {
	query := fmt.Sprintf("UPDATE %s SET username=$1 , password_hash=$2 WHERE id=$3", userTable)
	_, err := a.db.Exec(query, username, password, id)
	return err
}
