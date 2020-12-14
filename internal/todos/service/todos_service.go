package service

import (
	"github.com/mrsufgi/todo_api/internal/domain"
	log "github.com/sirupsen/logrus"
)

type todosService struct {
	tr domain.TodosRepository
}

func NewTodoService(tr domain.TodosRepository) domain.TodosService {
	log.Println("new logs")
	return &todosService{
		tr: tr,
	}
}
