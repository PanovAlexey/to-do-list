package services

import (
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/domain/entity/user"
)

type AuthorizationServiceInterface interface {
	repositories.AuthorizationRepositoryInterface
	CreateUser(user user.User) (int, error)
}
