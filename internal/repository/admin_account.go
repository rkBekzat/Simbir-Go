package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"vtb_api/internal/entities"
)

type admin struct {
	db *sqlx.DB
}

func NewAdmin(db *sqlx.DB) Admin {
	return &admin{db: db}
}

func (a *admin) GetAccounts(start, count int) ([]entities.User, error) {
	var users []entities.User
	query := fmt.Sprintf("SELECT TOP $2 * FROM %s WHERE id>=$1", userTable)
	err := a.db.Select(&users, query, start, count)
	return users, err
}

func (a *admin) GetAccountById(id int) (*entities.User, error) {
	var user entities.User
	query := fmt.Sprintf("SELECT username, name, is_admin, balace FROM %s WHERE id=$1", userTable)
	row := a.db.QueryRow(query, id)
	if err := row.Scan(&user.Username, &user.Name); err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *admin) CreateAccount(user *entities.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, is_admin, balance) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password, user.IsAdmin, user.Balance)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *admin) UpdateAccount(user *entities.User) error {
	query := fmt.Sprintf("UPDATE %s SET username=$1 , password_hash=$2, is_admin=$3, balance=$4 WHERE id=$5", userTable)
	_, err := a.db.Exec(query, user.Username, user.Password, user.IsAdmin, user.Balance, user.Id)
	return err
}

func (a *admin) DeleteAccount(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := a.db.Exec(query, id)
	return err
}
