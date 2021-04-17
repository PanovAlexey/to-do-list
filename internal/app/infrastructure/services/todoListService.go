package services

import (
	"todo-list/internal/app/application/repositories"
	"todo-list/internal/app/application/services"
)

type TodoListService struct {
	repository repositories.TodoListRepositoryInterface
}

func NewTodoListService(repository repositories.TodoListRepositoryInterface) services.TodoListServiceInterface {
	return &TodoListService{repository: repository}
}
