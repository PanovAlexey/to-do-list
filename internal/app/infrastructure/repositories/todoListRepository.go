package repositories

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/app/application/repositories"
)

type TodoListRepository struct {
}

func NewTodoListRepository(db *sqlx.DB) repositories.TodoListRepositoryInterface {
	return &TodoListRepository{}
}
