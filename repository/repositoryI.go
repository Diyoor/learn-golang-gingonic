package repository

import (
	"github.com/maxdev/go-gingonic/entity"
)

type RepositoryI interface {
	AddTodo(todo *entity.Todo) (int64, error)
	GetTodos() ([]entity.Todo, error)
	UpdateTodo(id int64, todo map[string]interface{}) (entity.Todo, error)
	DeleteTodo(id int64) (string, error)
	GetByID(id int64) (entity.Todo, error)
}
