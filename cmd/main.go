package main

import (
	todo_app "github.com/bakhodur-nazriev/todo-app"
	handler "github.com/bakhodur-nazriev/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo_app.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}