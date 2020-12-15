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

func (tr *pgTodosRepository) SearchTodos() (*[]domain.Todo, error) {
	query := "SELECT todo_id, done, name, details from todos"
	todo := &[]domain.Todo{}
	if err := tr.conn.Select(todo, query); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return todo, nil
}

func (tr *pgTodosRepository) ReadTodo(id int) (*domain.Todo, error) {
	query := "SELECT todo_id, done, name, details from todos WHERE todo_id = $1"
	todo := &domain.Todo{}
	if err := tr.conn.Get(todo, query, id); err != nil {
		log.Errorf("query error: %v", err)
		return nil, err
	}
	return todo, nil
}

func (tr *pgTodosRepository) CreateTodo(todo domain.Todo) (int, error) {
	query := `INSERT INTO todos (name, details) VALUES ($1, $2) RETURNING todo_id`
	var id int
	if err := tr.conn.QueryRow(query, todo.Name, todo.Details).Scan(&id); err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}
	return id, nil
}

func (tr *pgTodosRepository) UpdateTodo(id int, todo domain.Todo) (int64, error) {
	query := `UPDATE todos SET name=$2, details=$3, done=$4 WHERE todo_id=$1`
	res, err := tr.conn.Exec(query, id, todo.Name, todo.Details, todo.Done)

	if err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error checking affected rows: %v", err)
	}

	return rowsAffected, nil
}

func (tr *pgTodosRepository) DeleteTodo(id int) (int64, error) {
	query := `DELETE FROM todos WHERE todo_id = $1`
	res, err := tr.conn.Exec(query, id)

	if err != nil {
		log.Errorf("query error: %v", err)
		return -1, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("error checking affected rows: %v", err)
	}

	return rowsAffected, nil
}
