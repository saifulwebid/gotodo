package gotodo

// Service will implement this library's use cases.
type Service interface {
	Get(id int) (Todo, error)
	GetAll() []Todo
	GetPending() []Todo
	GetFinished() []Todo
	Add(title string, description *string) (Todo, error)
	Edit(todo Todo) (Todo, error)
	MarkAsDone(todo Todo) Todo
	Delete(todo Todo) error
	DeleteFinished()
}

type todolistService struct{}

func NewService() Service {
	return todolistService{}
}
