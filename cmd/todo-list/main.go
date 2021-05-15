package main

import (
	"github.com/spf13/viper"
	"log"
	"todo-list/internal/app/infrastructure/databases"
	"todo-list/internal/app/infrastructure/repositories"
	"todo-list/internal/app/infrastructure/services"
	"todo-list/internal/app/presentation"
	"todo-list/internal/app/servers"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())
	}

	database, err := databases.NewPostgresDB(databases.Config{
		Host:     viper.GetString("db_host"),
		Port:     viper.GetString("db_port"),
		Username: viper.GetString("db_username"),
		Password: viper.GetString("db_password"),
		DBName:   viper.GetString("db_name"),
		SSLMode:  viper.GetString("db_ssl_mode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize database: %s", err.Error())
	}

	authorizationRepository := repositories.NewAuthorizationRepository(database)
	todoItemRepository := repositories.NewTodoItemRepository(database)
	todoListRepository := repositories.NewTodoListRepository(database)

	authorizationService := services.NewAuthorizationService(authorizationRepository)
	todoItemService := services.NewTodoItemService(todoItemRepository)
	todoListService := services.NewTodoListService(todoListRepository)

	handler := presentation.NewHandler(authorizationService, todoItemService, todoListService)

	server := new(servers.Server)

	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
