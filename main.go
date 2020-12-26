package main

import (
	"log"
	"todo-list/internal/app/servers"
)

func main() {
	server := new(servers.Server)

	if err := server.Run("8000"); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
