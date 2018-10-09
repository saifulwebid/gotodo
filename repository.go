package gotodo

type Repository interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetWhere(status Status) []*Todo
	Insert(entityTodo *Todo) (*Todo, error)
	Update(entityTodo *Todo) error
	Delete(entityTodo *Todo) error
	DeleteWhere(status Status)
}
