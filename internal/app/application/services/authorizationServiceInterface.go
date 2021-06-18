package services

import (
	"todo-list/internal/app/domain/entity/user"
)

type AuthorizationServiceInterface interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
