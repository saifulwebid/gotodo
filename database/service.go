package database

import (
	"github.com/saifulwebid/gotodo"
)

type Service interface {
	Get(id int) (*gotodo.Todo, error)
	GetAll() []gotodo.Todo
	GetWhere(status gotodo.Status) []gotodo.Todo
	Insert(title string, description *string, status bool) (*gotodo.Todo, error)
	Update(todo gotodo.Todo) error
	Delete(todo gotodo.Todo) error
	DeleteWhere(status gotodo.Status) error
}
