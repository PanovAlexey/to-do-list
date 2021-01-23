package repositories

import (
	"todo-list/internal/app/application/repositories"
)

type TodoItemRepository struct {
}

func NewTodoItemRepository() repositories.TodoItemRepositoryInterface {
	return &TodoItemRepository{}
}
