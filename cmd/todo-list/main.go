package main

import (
	"log"
	"todo-list/internal/app/infrastructure/repositories"
	"todo-list/internal/app/infrastructure/services"
	"todo-list/internal/app/presentation"
	"todo-list/internal/app/servers"
)

func main() {
	authorizationRepository := repositories.NewAuthorizationRepository()
	todoItemRepository := repositories.NewTodoItemRepository()
	todoListRepository := repositories.NewTodoListRepository()

	authorizationService := services.NewAuthorizationService(authorizationRepository)
	todoItemService := services.NewTodoItemService(todoItemRepository)
	todoListService := services.NewTodoListService(todoListRepository)

	handler := presentation.NewHandler(authorizationService, todoItemService, todoListService)

	server := new(servers.Server)

	if err := server.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
