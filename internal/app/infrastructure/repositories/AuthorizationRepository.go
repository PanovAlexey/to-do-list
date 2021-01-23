package repositories

import (
	"todo-list/internal/app/application/repositories"
)

type AuthorizationRepository struct {
}

func NewAuthorizationRepository() repositories.AuthorizationRepositoryInterface {
	return &AuthorizationRepository{}
}
