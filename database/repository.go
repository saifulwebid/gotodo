package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/saifulwebid/gotodo"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

// NewRepository will return an implementation of a gotodo.Repository that
// connected to a MySQL database. It will return an error otherwise.
func NewRepository() (gotodo.Repository, error) {
	config := mysql.NewConfig()

	config.User = os.Getenv("GOTODO_DB_USER")
	config.Passwd = os.Getenv("GOTODO_DB_PASS")
	config.Net = "tcp"
	config.Addr = fmt.Sprint(os.Getenv("GOTODO_DB_HOST"), ":", os.Getenv("GOTODO_DB_PORT"))
	config.DBName = os.Getenv("GOTODO_DB_NAME")

	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&gotodo.Todo{})

	return &repository{db: db}, nil
}

// Get will return a Todo with supplied ID from the database, or return an error
// if a database error occured.
func (s *repository) Get(id int) (*gotodo.Todo, error) {
	todo := &gotodo.Todo{}

	res := s.db.First(&todo)
	if res.Error != nil {
		return nil, res.Error
	}

	return todo, nil
}

// GetAll will return all Todos in the database.
func (s *repository) GetAll() []*gotodo.Todo {
	todos := []*gotodo.Todo{}

	s.db.Find(&todos)

	return todos
}

// GetWhereDone will return all Todos with matching status.
func (s *repository) GetWhereDone(done bool) []*gotodo.Todo {
	todos := []*gotodo.Todo{}

	s.db.Where("done = ?", done).Find(&todos)

	return todos
}

// Insert will insert a Todo to the database and return the created Todo from
// the database in gotodo.Todo format, or return an error if a database error
// occured.
func (s *repository) Insert(todo *gotodo.Todo) (*gotodo.Todo, error) {
	res := s.db.Create(todo)
	if res.Error != nil {
		return nil, res.Error
	}

	return todo, nil
}

// Update will update the todo in the database, and return an error if a
// database error occured.
func (s *repository) Update(todo *gotodo.Todo) error {
	res := s.db.Save(todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Delete will delete the todo in the database, and return an error if a
// database error occured.
func (s *repository) Delete(todo *gotodo.Todo) error {
	if todo.ID == 0 {
		return errors.New("Invalid ID to delete")
	}

	res := s.db.Delete(&todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteWhereDone will delete all todos with matching status.
func (s *repository) DeleteWhereDone(done bool) {
	s.db.Where("done = ?", done).Find(gotodo.Todo{})
}
