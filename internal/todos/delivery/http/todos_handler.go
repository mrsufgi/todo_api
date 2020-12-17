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
	HTTPStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (e *ResponseError) WriteToResponse(w http.ResponseWriter) {
	w.WriteHeader(e.HTTPStatus)
	fmt.Fprint(w, e.ToJSON())
}

func (e *ResponseError) ToJSON() string {
	j, err := json.Marshal(e)
	if err != nil {
		return `{"code":50001,"message":"unable to marshal error"}`
	}
	return string(j)
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

	r.GET("/todos", handler.searchTodos)
	r.GET("/todos/:id", handler.readTodo)
	r.POST("/todos", handler.createTodo)
	r.PUT("/todos/:id", handler.updateTodo)
	r.DELETE("/todos/:id", handler.deleteTodo)

	return handler
}

func (p *TodosHandler) searchTodos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	rec, err := p.TService.SearchTodos()
	if err != nil {
		herr := &ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "unable to search todos"}
		herr.WriteToResponse(w)
		return
	}

	if err := json.NewEncoder(w).Encode(rec); err != nil {
		log.Errorf("unable to encode response %v", err)
	}
}

func (p *TodosHandler) readTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Errorf("unable to convert the string into int.  %v", err)
	}

	rec, err := p.TService.ReadTodo(id)
	if err != nil {
		log.Errorf("unable to read todo %v", err)
		herr := &ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to read todo"}
		herr.WriteToResponse(w)
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
		log.Errorf("unable to read todo %v", err)
		herr := &ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to update todo"}
		herr.WriteToResponse(w)
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
		log.Errorf("unable to update todo %v", err)
		herr := &ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to update todo"}
		herr.WriteToResponse(w)
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
		log.Errorf("unable to delete todo %v", err)
		herr := &ResponseError{HTTPStatus: http.StatusBadRequest, Code: 40001, Message: "Unable to delete todo"}
		herr.WriteToResponse(w)
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
