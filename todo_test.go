package gotodo

import (
	"testing"
)

func TestValidity(t *testing.T) {
	t.Run("valid cases", func(t *testing.T) {
		t.Parallel()

		t.Run("all attributes filled", func(t *testing.T) {
			t.Parallel()

			todo := Todo{1, "Test", "Test", false}

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("ID not filled", func(t *testing.T) {
			t.Parallel()

			todo := Todo{0, "Title", "Description", false}

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("description not filled", func(t *testing.T) {
			t.Parallel()

			todo := Todo{2, "Title", "", false}

			if !todo.isValid() {
				t.Fail()
			}
		})

		t.Run("status is Finished", func(t *testing.T) {
			t.Parallel()

			todo := Todo{3, "Title", "Description", true}

			if !todo.isValid() {
				t.Fail()
			}
		})
	})

	t.Run("invalid case", func(t *testing.T) {
		t.Parallel()

		t.Run("title is blank", func(t *testing.T) {
			todo := Todo{4, "", "Description", false}

			if todo.isValid() {
				t.Fail()
			}
		})
	})
}
