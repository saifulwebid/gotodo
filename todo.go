package gotodo

// Todo is a struct representation of, well, a todo.
type Todo struct {
	id     int
	status Status

	Title       string
	Description *string
}

// NewTodo will return a new todo based on the arguments.
func NewTodo(id int, title string, description *string, status Status) Todo {
	todo := Todo{
		id:          id,
		Title:       title,
		Description: description,
		status:      status,
	}

	return todo
}

// ID returns the todo's ID in the system.
func (t Todo) ID() int {
	return t.id
}

// Status returns the todo's status in the system.
func (t Todo) Status() Status {
	return t.status
}

// MarkAsDone will mark a todo as done.
func (t Todo) MarkAsDone() {
	t.status = Finished
}

// IsValid returns true if the todo is valid; false otherwise.
func (t Todo) IsValid() bool {
	if t.Title == "" {
		return false
	}

	return true
}
