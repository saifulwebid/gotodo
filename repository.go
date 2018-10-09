package gotodo

// Repository is an interface to store todos to a persistent storage.
type Repository interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetWhereDone(done bool) []*Todo
	Insert(todo *Todo) (*Todo, error)
	Update(todo *Todo) error
	Delete(todo *Todo) error
	DeleteWhereDone(done bool)
}
