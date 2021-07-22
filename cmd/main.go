package main

import (
	todo_app "github.com/bakhodur-nazriev/todo-app"
	handler "github.com/bakhodur-nazriev/todo-app/pkg/handler"
	"github.com/bakhodur-nazriev/todo-app/pkg/repository"
	"github.com/bakhodur-nazriev/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializtion configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
