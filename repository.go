package gotodo

// Repository is an interface to store todos to a persistent storage.
type Repository interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetWhere(status Status) []*Todo
	Insert(entityTodo *Todo) (*Todo, error)
	Update(entityTodo *Todo) error
	Delete(entityTodo *Todo) error
	DeleteWhere(status Status)
}
