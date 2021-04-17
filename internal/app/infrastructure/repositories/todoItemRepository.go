package repositories

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/app/application/repositories"
)

type TodoItemRepository struct {
}

//@ToDo: move sqlx.DB dependency to interface
func NewTodoItemRepository(db *sqlx.DB) repositories.TodoItemRepositoryInterface {
	return &TodoItemRepository{}
}
