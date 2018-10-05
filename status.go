package gotodo

// Status is a type representing a todo's status.
type Status int

const (
	// Pending is assigned to a todo's status if it is not done yet.
	Pending Status = iota

	// Finished is assigned to a todo's status if it is done.
	Finished
)
