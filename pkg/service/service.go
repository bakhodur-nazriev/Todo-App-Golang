package service

import (
	todo_app "github.com/bakhodur-nazriev/todo-app"
	"github.com/bakhodur-nazriev/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
