# gotodo

Gotodo is a Golang package that wraps database call to store and manage todos. Gotodo use MySQL to store all todos.

## Dependency

Gotodo needs a MySQL database to store its todos.

## Installation

Add this package to your Golang application.

```go
dep ensure -add github.com/saifulwebid/gotodo
```

### Configuration

Prepare an empty MySQL database in your environment. Then, copy `env.sample` to `.env` and change the variables based on your environment.

Gotodo will handle the migration automatically.

## Scope

Gotodo is able to:

* add a new todo,
* edit a pending todo,
* mark a pending todo as done,
* retrieve detail of a todo,
* list all todos,
* list pending todos,
* list finished todos,
* delete a pending todo,
* delete finished todos.

A todo can be either pending or finished. The status is determined by user actions.

A todo consists of:

* ID (integer, required),
* title (string, required),
* description (string, optional),
* status (enum, required, either `gotodo.Pending` or `gotodo.Finished`).

## Usage

Set up a service:

```go
import (
    "fmt"
    "os"

    "github.com/saifulwebid/gotodo"
    "github.com/saifulwebid/gotodo/database"
)

# In your entry point function
repo, err := database.NewRepository()
if err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
}

service := gotodo.NewService(repo)
```

### Add a new todo

```go
todo, err := service.Add("Buy battery", "Buy two AA batteries from Bukalapak")
if err != nil {
    // Todo is most likely invalid; a todo requires the title to be not an empty string
    fmt.Fprintln(os.Stderr, err.Error())
}
// todo will contain a Todo object from database.
// This todo is pending; it means that it is not done.
```

### Edit a pending todo

```go
todo.Title = "Buy two batteries"
err = service.Edit(todo)
if err != nil {
    // Todo is most likely invalid; a todo requires the title to be not an empty string
    fmt.Fprintln(os.Stderr, err.Error())
}
```

### Mark a pending todo as done

```go
err = service.MarkAsDone(todo)
if err != nil {
    // A database error has occured.
    fmt.Fprintln(os.Stderr, err.Error())
}
```

### Retrieve detail of a todo

```go
todo, err = service.Get(1)
if err != nil {
    // A Todo with ID = 1 is not found.
    fmt.Fprintln(os.Stderr, err.Error())
}
```

### List all todos

```go
todos := service.GetAll()
```

### List pending todos

```go
todos = service.GetPending()
```

### List finished todos

```go
todos = service.GetFinished()
```

### Delete a pending todo

```go
error = service.Delete(todo)
if err != nil {
    // That Todo is finished already (and should be deleted using
    // service.DeleteFinished() method).
    fmt.Fprintln(os.Stderr, err.Error())
}
```

### Delete finished todos

```go
service.DeleteFinished()
```
