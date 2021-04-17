package services

import "todo-list/internal/app/application/repositories"

type TodoListServiceInterface interface {
	repositories.TodoListRepositoryInterface
}
