package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/domain/entity/user"
	"todo-list/internal/app/infrastructure/databases"
)

type AuthorizationRepository struct {
	db *sqlx.DB
}

// @ToDo: hide sqlx package by interface
func NewAuthorizationRepository(db *sqlx.DB) repositories.AuthorizationRepositoryInterface {
	return &AuthorizationRepository{db: db}
}

func (repository AuthorizationRepository) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", databases.UsersTable)
	row := repository.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository AuthorizationRepository) GetUser(username, password string) (user.User, error) {
	var user user.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", databases.UsersTable)
	err := repository.db.Get(&user, query, username, password)

	return user, err
}
