package repositories

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/app/application/repositories"
)

type AuthorizationRepository struct {
}

func NewAuthorizationRepository(db *sqlx.DB) repositories.AuthorizationRepositoryInterface {
	return &AuthorizationRepository{}
}
