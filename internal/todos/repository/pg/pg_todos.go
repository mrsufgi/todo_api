package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrsufgi/todo_api/internal/domain"
	log "github.com/sirupsen/logrus"
)

type pgTodosRepository struct {
	conn *sqlx.DB
}

// NewPgTodosRepository will create an object that represent the X interface
func NewPgTodosRepository(conn *sqlx.DB) domain.TodosRepository {
	r := &pgTodosRepository{
		conn: conn,
	}

	return r
}

func (tr *pgTodosRepository) GetTodo(id int) (*domain.Todo, error) {
	query := "SELECT todo_id, done, name, details from todos WHERE todo_id = $1"
	todo := &domain.Todo{}
	if err := tr.conn.Get(todo, query, id); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return todo, nil
}
