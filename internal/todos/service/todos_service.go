package service

import (
	"github.com/mrsufgi/todo_api/internal/domain"
	log "github.com/sirupsen/logrus"
)

type todosService struct {
	tr domain.TodosRepository
}

func NewTodoService(tr domain.TodosRepository) domain.TodosService {
	return &todosService{
		tr: tr,
	}
}

// TODO: add search variables
func (ts *todosService) SearchTodos() (*[]domain.Todo, error) {
	res, err := ts.tr.SearchTodos()
	if len(*res) == 0 {
		log.Info("unable to search todos")
	}
	return res, err
}

func (ts *todosService) CreateTodo(todo domain.Todo) (int, error) {
	res, err := ts.tr.CreateTodo(todo)
	if res == -1 {
		log.Infof("unable to create todo: %v", todo)
		return -1, err
	}
	return res, nil
}

func (ts *todosService) ReadTodo(id int) (*domain.Todo, error) {
	res, err := ts.tr.ReadTodo(id)
	if res == nil {
		log.Infof("unable to find todo: %v", id)
		return nil, err
	}
	return res, nil
}

func (ts *todosService) UpdateTodo(id int, todo domain.Todo) error {
	res, err := ts.tr.UpdateTodo(id, todo)
	if res != 1 {
		log.Infof("unable to update todo: %v, %v", id, todo)
	}
	return err
}

func (ts *todosService) DeleteTodo(id int) error {
	res, err := ts.tr.DeleteTodo(id)
	if res != 1 {
		log.Infof("unable to delete todo: %v", id)
	}
	return err
}
