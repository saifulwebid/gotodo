package database

import (
	"github.com/saifulwebid/gotodo"
)

type Todo struct {
	ID          int    `gorm:"PRIMARY_KEY"`
	Title       string `gorm:"NOT NULL"`
	Description *string
	Done        *bool `gorm:"NOT NULL"`
}

func fromEntity(todo *gotodo.Todo) Todo {
	var done bool
	if todo.Status() == gotodo.Finished {
		done = true
	} else {
		done = false
	}

	return Todo{
		ID:          todo.ID(),
		Title:       todo.Title,
		Description: todo.Description,
		Done:        &done,
	}
}

func (t Todo) asEntity() *gotodo.Todo {
	var status gotodo.Status
	if *t.Done {
		status = gotodo.Finished
	} else {
		status = gotodo.Pending
	}

	return gotodo.NewTodo(
		t.ID,
		t.Title,
		t.Description,
		status,
	)
}
