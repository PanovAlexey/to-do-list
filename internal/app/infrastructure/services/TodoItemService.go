package services

import (
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/application/services"
)

type TodoItemService struct {
	repository repositories.TodoItemRepositoryInterface
}

func NewTodoItemService(repository repositories.TodoItemRepositoryInterface) services.TodoItemServiceInterface {
	return &TodoItemService{repository: repository}
}
