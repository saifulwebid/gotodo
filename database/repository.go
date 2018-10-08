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

func (s *repository) Get(id int) (*gotodo.Todo, error) {
	todo := Todo{}

	res := s.db.First(&todo)
	if res.Error != nil {
		return nil, res.Error
	}

	entity := todo.AsEntity()
	return &entity, nil
}

func (s *repository) getEntities(populate func(*[]Todo)) []gotodo.Todo {
	todos := []Todo{}

	populate(&todos)

	ret := []gotodo.Todo{}
	for _, el := range todos {
		ret = append(ret, el.AsEntity())
	}

	return ret
}

func (s *repository) GetAll() []gotodo.Todo {
	return s.getEntities(func(todos *[]Todo) {
		s.db.Find(&todos)
	})
}

func (s *repository) GetWhere(status gotodo.Status) []gotodo.Todo {
	done := (status == gotodo.Finished)

	return s.getEntities(func(todos *[]Todo) {
		s.db.Where(&Todo{Done: &done}).Find(&todos)
	})
}

func (s *repository) Insert(title string, description *string, done bool) (*gotodo.Todo, error) {
	todo := Todo{
		Title:       title,
		Description: description,
		Done:        &done,
	}

	res := s.db.Create(&todo)
	if res.Error != nil {
		return nil, res.Error
	}

	entity := todo.AsEntity()
	return &entity, nil
}

func (s *repository) Update(entityTodo gotodo.Todo) error {
	todo := FromEntity(entityTodo)

	res := s.db.Save(todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *repository) Delete(entityTodo gotodo.Todo) error {
	if entityTodo.ID() == 0 {
		return errors.New("Invalid ID to delete")
	}
	todo := FromEntity(entityTodo)

	res := s.db.Delete(&todo)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *repository) DeleteWhere(status gotodo.Status) {
	done := (status == gotodo.Finished)
	filter := Todo{Done: &done}

	s.db.Delete(filter)
}
