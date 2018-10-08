package database

import (
	"github.com/saifulwebid/gotodo"
)

type Service interface {
	Get(id int) (*gotodo.Todo, error)
	GetAll() []gotodo.Todo
	GetWhere(status gotodo.Status) []gotodo.Todo
	Insert(title string, description *string) (*gotodo.Todo, error)
	Update(todo gotodo.Status) error
	Delete(todo gotodo.Status) error
	DeleteWhere(status gotodo.Status) error
}
