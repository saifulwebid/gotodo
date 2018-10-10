package gotodo

// Todo is a struct representation of, well, a todo.
type Todo struct {
	ID          int    `gorm:"PRIMARY_KEY" json:"id"`
	Title       string `gorm:"NOT NULL" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"NOT NULL" json:"done"`
}

// IsValid returns true if the todo is valid; false otherwise.
func (t *Todo) isValid() bool {
	return t.Title != ""
}
