package services

import "todo-list/internal/app/application/repositories"

type AuthorizationServiceInterface interface {
	repositories.AuthorizationRepositoryInterface
}
