package http

import (
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/todo_api/internal/domain"
	log "github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

//
type TodosHandler struct {
	TService domain.TodosService
}

//
func NewTodoHandler(r *httprouter.Router, ts domain.TodosService) *TodosHandler {
	handler := &TodosHandler{
		TService: ts,
	}

	r.GET("/todos/user/:id", handler.userTodo) // TODO: design on endpoints
	return handler
}

func (p *TodosHandler) userTodo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Debug("user todo")

	_, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("failed to parse body", err.Error())
		return
	}

	// id := ps.ByName("id")
	//
	/* rec, err := p.TService.GenerateUserTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf("Unable to generate user todo %v", err)
		return
	}

	res, err := json.Marshal(rec)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf("Unable marshal result: %v", err)
		return
	} */
	w.Header().Set("Content-Type", "application/json")
	// w.Write(res)
}
