package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/todo_api/internal/domain"
	log "github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type ResponseMessage struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type TodosHandler struct {
	TService domain.TodosService
}

func NewTodosHandler(r *httprouter.Router, ts domain.TodosService) *TodosHandler {
	handler := &TodosHandler{
		TService: ts,
	}

	r.GET("/todos/:id", handler.readTodo)
	r.POST("/todos", handler.createTodo)
	r.PUT("/todos/:id", handler.updateTodo)
	r.DELETE("/todos/:id", handler.deleteTodo)

	return handler
}

func (p *TodosHandler) readTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	rec, err := p.TService.ReadTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Errorf("unable to read todo %v", err)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *TodosHandler) createTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var todo domain.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Errorf("unable to parse the body.  %v", err)
	}

	id, err := p.TService.CreateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Errorf("unable to read todo %v", err)
		return
	}

	res := ResponseMessage{
		ID:      int64(id),
		Message: "Todo created successfully",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *TodosHandler) updateTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	var todo domain.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Errorf("unable to parse the body.  %v", err)
	}

	affected, err := p.TService.UpdateTodo(id, todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Errorf("unable to update todo %v", err)
		return
	}

	msg := fmt.Sprintf("todo updated successfully. total rows affected %v", affected)

	res := ResponseMessage{
		ID:      int64(id),
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *TodosHandler) deleteTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	affected, err := p.TService.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Errorf("unable to delete todo %v", err)
		return
	}

	msg := fmt.Sprintf("todo deleted successfully. total rows affected %v", affected)

	res := ResponseMessage{
		ID:      int64(id),
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}
