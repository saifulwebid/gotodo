package gotodo

// Service will implement this library's use cases.
type Service interface {
	Get(id int) Todo
	GetAll() []Todo
	GetPending() []Todo
	GetFinished() []Todo
	Add(title string, description *string) Todo
	Edit(todo Todo) Todo
	MarkAsDone(todo Todo) Todo
	Delete(todo Todo)
	DeleteFinished()
}

type todolistService struct{}

func NewService() Service {
	return todolistService{}
}
