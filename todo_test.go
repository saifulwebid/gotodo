package gotodo

import (
	"testing"
)

func TestMarkAsDone(t *testing.T) {
	todo := NewTodo(1, "Title", "Description", Pending)

	t.Run("Pending => Finished", func(t *testing.T) {
		todo.markAsDone()
		if todo.Status() != Finished {
			t.FailNow()
		}
	})

	t.Run("Finished => Finished", func(t *testing.T) {
		todo.markAsDone()
		if todo.Status() != Finished {
			t.FailNow()
		}
	})
}

func TestValidity(t *testing.T) {
	t.Run("valid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("all attributes filled", func(t *testing.T) {
			t.Parallel()

			todo := NewTodo(1, "Test", "Test", Pending)

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("ID not filled", func(t *testing.T) {
			t.Parallel()

			todo := NewTodo(0, "Title", "Description", Pending)

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("description not filled", func(t *testing.T) {
			t.Parallel()

			todo := NewTodo(2, "Title", "", Pending)

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("status is Finished", func(t *testing.T) {
			t.Parallel()

			todo := NewTodo(3, "Title", "Description", Finished)

			if !todo.isValid() {
				t.Fail()
			}
		})
	})

	t.Run("invalid case", func(t *testing.T) {
		t.Parallel()

		t.Run("title is blank", func(t *testing.T) {
			todo := NewTodo(4, "", "Description", Pending)

			if todo.isValid() {
				t.Fail()
			}
		})
	})
}
