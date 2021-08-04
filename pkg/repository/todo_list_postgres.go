package repository

import (
	"fmt"
	todo_app "github.com/bakhodur-nazriev/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo_app.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo_app.TodoList, error) {
	var lists []todo_app.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.td = ul.list_id WHERE ul.user.id = $1", todoListsTable, usersListTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo_app.TodoList, error) {
	var list todo_app.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s 
								tl INNER JOIN %s ul on tl.td = ul.list_id WHERE ul.user.id = $1 AND ul.list_id = $2`, todoListsTable, usersListTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
