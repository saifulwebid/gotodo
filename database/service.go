package database

import (
	"github.com/saifulwebid/gotodo"
)

type Service interface {
	Get(id int) (*gotodo.Todo, error)
	GetAll() []gotodo.Todo
	GetWhere(status gotodo.Status) []gotodo.Todo
	Insert(title string, description *string, done bool) (*gotodo.Todo, error)
	Update(entityTodo gotodo.Todo) error
	Delete(entityTodo gotodo.Todo) error
	DeleteWhere(status gotodo.Status)
}
