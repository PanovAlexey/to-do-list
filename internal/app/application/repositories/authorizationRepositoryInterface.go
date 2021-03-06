package repositories

import "todo-list/internal/app/domain/entity/user"

type AuthorizationRepositoryInterface interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}
