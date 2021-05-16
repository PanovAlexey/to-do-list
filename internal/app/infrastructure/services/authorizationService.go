package services

import (
	"crypto/sha1"
	"fmt"
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/application/services"
	"todo-list/internal/app/domain/entity/user"
)

const salt = "wertyuiopasdfghjkl"

type AuthorizationService struct {
	repository repositories.AuthorizationRepositoryInterface
}

func (service *AuthorizationService) CreateUser(user user.User) (int, error) {
	user.Password = service.generatePasswordHash(user.Password)
	return service.repository.CreateUser(user)
}

func NewAuthorizationService(repository repositories.AuthorizationRepositoryInterface) services.AuthorizationServiceInterface {
	return &AuthorizationService{repository: repository}
}

func (service *AuthorizationService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
