package domain

type Todo struct {
	ID      int    `db:"todo_id" json:"id,omitempty"`
	Done    bool   `db:"done" json:"done"` // note: consider using a pointer (if u want omitempty)
	Name    string `db:"name" json:"name,omitempty"`
	Details string `db:"details" json:"details,omitempty"`
}

//go:generate mockgen -destination=mocks/mock_todos_repository.go -package=mocks . TodosRepository
type TodosRepository interface {
	SearchTodos() (*[]Todo, error)
	CreateTodo(todo Todo) (int, error)
	ReadTodo(id int) (*Todo, error)
	UpdateTodo(id int, todo Todo) (int64, error)
	DeleteTodo(id int) (int64, error) // TODO: todo archive (soft delete)
}

//go:generate mockgen -destination=mocks/mock_todos_service.go -package=mocks . TodosService
type TodosService interface {
	SearchTodos() (*[]Todo, error)
	CreateTodo(todo Todo) (int, error)
	ReadTodo(id int) (*Todo, error)
	UpdateTodo(id int, todo Todo) (int64, error)
	DeleteTodo(id int) (int64, error)
}
