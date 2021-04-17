package services

import "todo-list/internal/app/application/repositories"

type TodoItemServiceInterface interface {
	repositories.TodoItemRepositoryInterface
}
