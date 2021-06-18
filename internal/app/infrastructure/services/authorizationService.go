package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/application/services"
	"todo-list/internal/app/domain/entity/user"
)

const (
	salt       = "wertyuiopasdfghjkl"
	tokenTTL   = 12 * time.Hour
	signingKey = "qweqr78939870424&(#$@"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthorizationService struct {
	repository repositories.AuthorizationRepositoryInterface
}

func NewAuthorizationService(repository repositories.AuthorizationRepositoryInterface) services.AuthorizationServiceInterface {
	return &AuthorizationService{repository: repository}
}

func GetUser(username, password string) (user.User, error) {
	return user.User{}, nil
}

func (service *AuthorizationService) CreateUser(user user.User) (int, error) {
	user.Password = service.generatePasswordHash(user.Password)
	return service.repository.CreateUser(user)
}

func (service *AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := service.repository.GetUser(username, service.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (service *AuthorizationService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (service *AuthorizationService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
