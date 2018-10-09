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

	db.AutoMigrate(&Todo{})

	return &repository{db: db}, nil
}

// Get will return a Todo with supplied ID from the database, or return an error
// if a database error occured.
func (s *repository) Get(id int) (*gotodo.Todo, error) {
	todo := Todo{}

	res := s.db.First(&todo)
	if res.Error != nil {
		return nil, res.Error
	}

	entity := todo.asEntity()
	return entity, nil
}

func (s *repository) getEntities(populate func(*[]Todo)) []*gotodo.Todo {
	todos := []Todo{}

	populate(&todos)

	ret := []*gotodo.Todo{}
	for _, el := range todos {
		ret = append(ret, el.asEntity())
	}

	return ret
}

// GetAll will return all Todos in the database.
func (s *repository) GetAll() []*gotodo.Todo {
	return s.getEntities(func(todos *[]Todo) {
		s.db.Find(&todos)
	})
}

// GetWhere will return all Todos with matching status.
func (s *repository) GetWhere(status gotodo.Status) []*gotodo.Todo {
	done := (status == gotodo.Finished)

	return s.getEntities(func(todos *[]Todo) {
		s.db.Where(&Todo{Done: &done}).Find(&todos)
	})
}

// Insert will insert a Todo to the database and return the created Todo from
// the database in gotodo.Todo format, or return an error if a database error
// occured.
func (s *repository) Insert(entityTodo *gotodo.Todo) (*gotodo.Todo, error) {
	todo := fromEntity(entityTodo)

	res := s.db.Create(&todo)
	if res.Error != nil {
		return nil, res.Error
	}

	entity := todo.asEntity()
	return entity, nil
}

// Update will update the todo in the database, and return an error if a
// database error occured.
func (s *repository) Update(entityTodo *gotodo.Todo) error {
	todo := fromEntity(entityTodo)

	res := s.db.Save(todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Delete will delete the todo in the database, and return an error if a
// database error occured.
func (s *repository) Delete(entityTodo *gotodo.Todo) error {
	if entityTodo.ID() == 0 {
		return errors.New("Invalid ID to delete")
	}
	todo := fromEntity(entityTodo)

	res := s.db.Delete(&todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// DeleteWhere will delete all todos with matching status.
func (s *repository) DeleteWhere(status gotodo.Status) {
	done := (status == gotodo.Finished)
	filter := Todo{Done: &done}

	s.db.Delete(filter)
}
