package gotodo

type Repository interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetWhere(status Status) []*Todo
	Insert(title string, description string, done bool) (*Todo, error)
	Update(entityTodo *Todo) error
	Delete(entityTodo *Todo) error
	DeleteWhere(status Status)
}
