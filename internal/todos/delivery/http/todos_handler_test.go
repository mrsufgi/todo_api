package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/mrsufgi/todo_api/internal/domain"
	"github.com/mrsufgi/todo_api/internal/domain/mocks"
	tdh "github.com/mrsufgi/todo_api/internal/todos/delivery/http"
)

func TestNewTodosHandler(t *testing.T) {
	type args struct {
		r  *httprouter.Router
		ts domain.TodosService
	}
	tests := []struct {
		name string
		args args
		want *tdh.TodosHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tdh.NewTodosHandler(tt.args.r, tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodosHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createReadRequest(t *testing.T, id int) *http.Request {
	r, err := http.NewRequest("GET", fmt.Sprintf("/todos/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	return r
}
func TestTodosHandler_readTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := httprouter.New()
	s := mocks.NewMockTodosService(ctrl)
	w := httptest.NewRecorder()
	r := createReadRequest(t, 1)

	type fields struct {
		TService domain.TodosService
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"happy read todo call", fields{TService: s}, args{w: w, r: r}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tdh.NewTodosHandler(router, tt.fields.TService)
			s.EXPECT().ReadTodo(gomock.Any())
			router.ServeHTTP(tt.args.w, tt.args.r)
			resp := w.Result()
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected status code %d", resp.StatusCode)
			}
		})
	}
}

// TODO: add missing API tests
