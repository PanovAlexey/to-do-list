package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"todo-list/internal/app/infrastructure/databases"
	"todo-list/internal/app/infrastructure/repositories"
	"todo-list/internal/app/infrastructure/services"
	"todo-list/internal/app/presentation"
	"todo-list/internal/app/servers"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	database, err := databases.NewPostgresDB(databases.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err.Error())
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
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
