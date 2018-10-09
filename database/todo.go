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
	done := (todo.Status() == gotodo.Finished)

	var desc *string
	if todo.Description == "" {
		desc = nil
	} else {
		desc = &todo.Description
	}

	return Todo{
		ID:          todo.ID(),
		Title:       todo.Title,
		Description: desc,
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

	var desc string
	if t.Description == nil {
		desc = ""
	} else {
		desc = *t.Description
	}

	return gotodo.NewTodo(
		t.ID,
		t.Title,
		desc,
		status,
	)
}
