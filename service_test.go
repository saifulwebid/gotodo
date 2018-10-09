package gotodo

import (
	"errors"
	"testing"
)

func createRepo() *mockRepo {
	repo := mockRepo{1, map[int]*Todo{}}

	/* Insert sample data */
	repo.Insert(&Todo{Title: "Title 1", Description: "Description", Done: false})
	repo.Insert(&Todo{Title: "Title 2", Description: "Description", Done: true})
	repo.Insert(&Todo{Title: "Title 3", Description: "Description", Done: false})

	return &repo
}

func TestGet(t *testing.T) {
	service := NewService(createRepo())

	t.Run("get existing todo", func(t *testing.T) {
		_, err := service.Get(3)
		if err != nil {
			t.FailNow()
		}
	})

	t.Run("get non-existing todo", func(t *testing.T) {
		_, err := service.Get(5)
		if err == nil {
			t.FailNow()
		}
	})
}

func TestGetAll(t *testing.T) {
	service := NewService(createRepo())

	res := service.GetAll()

	if len(res) != 3 {
		t.FailNow()
	}
}

func TestGetPending(t *testing.T) {
	service := NewService(createRepo())

	res := service.GetPending()

	if len(res) != 2 {
		t.FailNow()
	}
}

func TestGetFinished(t *testing.T) {
	service := NewService(createRepo())

	res := service.GetFinished()

	if len(res) != 1 {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	service := NewService(createRepo())

	t.Run("valid todo", func(t *testing.T) {
		todosCount := len(service.GetAll())

		_, err := service.Add("valid todo", "description")

		if err != nil {
			t.FailNow()
		}

		if len(service.GetAll()) == todosCount {
			t.FailNow()
		}
	})

	t.Run("invalid todo", func(t *testing.T) {
		todosCount := len(service.GetAll())

		_, err := service.Add("", "description")

		if err == nil {
			t.FailNow()
		}

		if len(service.GetAll()) != todosCount {
			t.FailNow()
		}
	})
}

func TestMarkAsDone(t *testing.T) {
	service := NewService(createRepo())

	todo, _ := service.Get(1)
	service.MarkAsDone(todo)

	if !todo.Done {
		t.FailNow()
	}

	todo, _ = service.Get(1)
	if !todo.Done {
		t.FailNow()
	}
}

func TestDelete(t *testing.T) {
	service := NewService(createRepo())

	t.Run("delete pending todo", func(t *testing.T) {
		todosCount := len(service.GetAll())

		todo, err := service.Add("Test pending", "")
		if err != nil {
			t.Fatal("cannot create todo")
		}

		err = service.Delete(todo)
		if err != nil {
			t.FailNow()
		}

		if len(service.GetAll()) != todosCount {
			t.FailNow()
		}
	})

	t.Run("delete finished todo", func(t *testing.T) {
		todosCount := len(service.GetAll())

		todo, err := service.Add("Test finished", "")
		if err != nil {
			t.FailNow()
		}

		service.MarkAsDone(todo)

		err = service.Delete(todo)
		if err == nil {
			t.FailNow()
		}

		if len(service.GetAll()) == todosCount {
			t.FailNow()
		}
	})
}

func TestDeleteFinished(t *testing.T) {
	service := NewService(createRepo())

	service.DeleteFinished()

	if len(service.GetFinished()) > 0 {
		t.FailNow()
	}

	if len(service.GetPending()) == 0 {
		t.FailNow()
	}
}

type mockRepo struct {
	nextID  int
	storage map[int]*Todo
}

func (m *mockRepo) Get(id int) (*Todo, error) {
	for _, todo := range m.storage {
		if todo.ID == id {
			return todo, nil
		}
	}

	return nil, errors.New("not found")
}

func (m *mockRepo) GetAll() []*Todo {
	ret := []*Todo{}

	for _, todo := range m.storage {
		ret = append(ret, todo)
	}

	return ret
}

func (m *mockRepo) GetWhereDone(done bool) []*Todo {
	ret := []*Todo{}

	for _, todo := range m.storage {
		if todo.Done == done {
			ret = append(ret, todo)
		}
	}

	return ret
}

func (m *mockRepo) Insert(todo *Todo) (*Todo, error) {
	todo.ID = m.nextID
	m.nextID++

	m.storage[todo.ID] = todo

	return todo, nil
}

func (m *mockRepo) Update(todo *Todo) error {
	if _, exists := m.storage[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	m.storage[todo.ID].Title = todo.Title
	m.storage[todo.ID].Description = todo.Description
	m.storage[todo.ID].Done = todo.Done

	return nil
}

func (m *mockRepo) Delete(todo *Todo) error {
	delete(m.storage, todo.ID)

	return nil
}

func (m *mockRepo) DeleteWhereDone(done bool) {
	for id, todo := range m.storage {
		if todo.Done == done {
			delete(m.storage, id)
		}
	}
}
