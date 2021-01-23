package repositories

import (
	"todo-list/internal/app/application/repositories"
)

type TodoListRepository struct {
}

func NewTodoListRepository() repositories.TodoListRepositoryInterface {
	return &TodoListRepository{}
}
