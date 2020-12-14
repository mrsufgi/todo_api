package domain

type Todo struct {
	ID      int    `db:"todo_id"`
	Done    bool   `db:"done"`
	Name    string `db:"name"`
	Details string `db:"details"`
}

//go:generate mockgen -destination=mocks/mock_todos_repository.go -package=mocks . TodosRepository
type TodosRepository interface {
	GetTodo(id int) (*Todo, error)
}

//go:generate mockgen -destination=mocks/mock_todos_service.go -package=mocks . TodosService
type TodosService interface {
}
