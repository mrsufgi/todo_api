package service_test

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrsufgi/todo_api/internal/domain"
	"github.com/mrsufgi/todo_api/internal/domain/mocks"
	"github.com/mrsufgi/todo_api/internal/todos/service"
)

func String(x string) *string {
	return &x
}

func TestNewTodoService(t *testing.T) {
	type args struct {
		tr domain.TodosRepository
	}
	tests := []struct {
		name string
		args args
		want domain.TodosService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewTodoService(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodoService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todosService_SearchTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockTodosRepository(ctrl)

	type fields struct {
		tr domain.TodosRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]domain.Todo
		wantErr bool
	}{
		{"happy search todos", fields{tr: tr}, &[]domain.Todo{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewTodoService(
				tt.fields.tr,
			)
			tr.EXPECT().SearchTodos().Return(&[]domain.Todo{}, nil)

			got, err := ts.SearchTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("todosService.SearchTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todosService.SearchTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todosService_CreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockTodosRepository(ctrl)

	type fields struct {
		tr domain.TodosRepository
	}
	type args struct {
		todo domain.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy create todo", fields{tr: tr}, args{domain.Todo{ID: 0, Done: false, Name: String("Test"), Details: String("None")}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewTodoService(
				tt.fields.tr,
			)
			tr.EXPECT().CreateTodo(tt.args.todo).Return(tt.args.todo.ID, nil)

			got, err := ts.CreateTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("todosService.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("todosService.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todosService_ReadTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockTodosRepository(ctrl)

	type fields struct {
		tr domain.TodosRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Todo
		wantErr bool
	}{
		{"happy read todo", fields{tr: tr}, args{id: 0},
			&domain.Todo{ID: 0, Done: false, Name: String("Test"), Details: String("None")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewTodoService(
				tt.fields.tr,
			)

			// note: returning the 'tt.want', simplify the fake data and validation checks
			// the func doesn't alter the result.
			tr.EXPECT().ReadTodo(tt.args.id).Return(tt.want, nil)
			got, err := ts.ReadTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todosService.ReadTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todosService.ReadTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todosService_UpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockTodosRepository(ctrl)

	type fields struct {
		tr domain.TodosRepository
	}
	type args struct {
		id   int
		todo domain.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy update todo", fields{tr: tr}, args{id: 0, todo: domain.Todo{Done: true}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewTodoService(
				tt.fields.tr,
			)
			tr.EXPECT().UpdateTodo(tt.args.id, tt.args.todo).Return(int64(1), nil)
			got, err := ts.UpdateTodo(tt.args.id, tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("todosService.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todosService.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todosService_DeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tr := mocks.NewMockTodosRepository(ctrl)

	type fields struct {
		tr domain.TodosRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy delete todo", fields{tr: tr}, args{id: 0}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := service.NewTodoService(
				tt.fields.tr,
			)
			tr.EXPECT().DeleteTodo(tt.args.id).Return(int64(1), nil)
			got, err := ts.DeleteTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("todosService.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todosService.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
