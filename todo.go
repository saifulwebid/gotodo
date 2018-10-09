package gotodo

// Todo is a struct representation of, well, a todo.
type Todo struct {
	ID          int    `gorm:"PRIMARY_KEY"`
	Title       string `gorm:"NOT NULL"`
	Description string
	Done        bool `gorm:"NOT NULL"`
}

// IsValid returns true if the todo is valid; false otherwise.
func (t *Todo) isValid() bool {
	return t.Title != ""
}
