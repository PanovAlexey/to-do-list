package main

import (
	"log"
	"todo-list/internal/app/presentation"
	"todo-list/internal/app/servers"
)

func main() {
	handlers := new(presentation.Handler)
	server := new(servers.Server)

	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
