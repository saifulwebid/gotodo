package gotodo

// Service will implement this library's use cases.
type Service interface {
	Get(id int) (*Todo, error)
	GetAll() []*Todo
	GetPending() []*Todo
	GetFinished() []*Todo
	Add(title string, description *string) (*Todo, error)
	Edit(todo *Todo) error
	MarkAsDone(todo *Todo)
	Delete(todo *Todo) error
	DeleteFinished()
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(id int) (*Todo, error) {
	return s.repo.Get(id)
}

func (s *service) GetAll() []*Todo {
	return s.repo.GetAll()
}

func (s *service) GetPending() []*Todo {
	return s.repo.GetWhere(Pending)
}

func (s *service) GetFinished() []*Todo {
	return s.repo.GetWhere(Finished)
}

func (s *service) Add(title string, description *string) (*Todo, error) {
	return s.repo.Insert(title, description, false)
}

func (s *service) Edit(todo *Todo) error {
	return s.repo.Update(todo)
}

func (s *service) MarkAsDone(todo *Todo) {
	todo.markAsDone()
}

func (s *service) Delete(todo *Todo) error {
	return s.repo.Delete(todo)
}

func (s *service) DeleteFinished() {
	s.repo.DeleteWhere(Finished)
}
