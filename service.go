package gotodo

import (
	"errors"
)

// Service will implement this library's use cases.
type Service interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetPending() []*Todo
	GetFinished() []*Todo
	Add(title string, description string) (*Todo, error)
	Edit(todo *Todo) error
	MarkAsDone(todo *Todo) error
	Delete(todo *Todo) error
	DeleteFinished()
}

type service struct {
	repo Repository
}

// NewService will return an implementation of a Service. It should be supplied
// with an implementation of a Repository.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Get will return a Todo if it exists in the repository, and will return an
// error otherwise.
func (s *service) Get(id int) (*Todo, error) {
	return s.repo.Get(id)
}

// GetAll will return a collection of all Todos in the repository.
func (s *service) GetAll() []*Todo {
	return s.repo.GetAll()
}

// GetPending will return a collection of all pending Todos in the repository.
func (s *service) GetPending() []*Todo {
	return s.repo.GetWhereDone(false)
}

// GetFinished will return a collection of all finished Todos in the repository.
func (s *service) GetFinished() []*Todo {
	return s.repo.GetWhereDone(true)
}

// Add will create a pending Todo with supplied title and description.
// It will return a Todo created in the repository, or an error if the Todo
// is invalid (it means that you have supplied an empty title).
func (s *service) Add(title string, description string) (*Todo, error) {
	todo := Todo{
		Title:       title,
		Description: description,
		Done:        false,
	}

	if !todo.isValid() {
		return nil, errors.New("Todo is invalid")
	}

	return s.repo.Insert(&todo)
}

// Edit will update the Todo in the repository with values supplied in the
// todo object parameter. It also will return an error if the todo become
// invalid.
func (s *service) Edit(todo *Todo) error {
	if !todo.isValid() {
		return errors.New("Todo is invalid")
	}

	return s.repo.Update(todo)
}

// MarkAsDone will mark the todo as done and update it in the repository.
func (s *service) MarkAsDone(todo *Todo) error {
	todo.Done = true

	return s.repo.Update(todo)
}

// Delete will delete a Todo from the repository.
func (s *service) Delete(todo *Todo) error {
	if todo.Done {
		return errors.New("todo is already finished; delete using service.DeleteFinished()")
	}

	return s.repo.Delete(todo)
}

// DeleteFinished will delete all finished Todos in the repository.
func (s *service) DeleteFinished() {
	s.repo.DeleteWhereDone(true)
}
