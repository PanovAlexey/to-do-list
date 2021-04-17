package services

import (
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/application/services"
)

type AuthorizationService struct {
	repository repositories.AuthorizationRepositoryInterface
}

func NewAuthorizationService(repository repositories.AuthorizationRepositoryInterface) services.AuthorizationServiceInterface {
	return &AuthorizationService{repository: repository}
}
