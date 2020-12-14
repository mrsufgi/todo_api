package repository_test

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrsufgi/todo_api/internal/domain"
	repository "github.com/mrsufgi/todo_api/internal/todos/repository/pg"
	log "github.com/sirupsen/logrus"
)

func TestNewPgTodosRepository(t *testing.T) {
	type args struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want domain.TodosRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewPgTodosRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPgTodosRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgTodosRepository_GetTodo(t *testing.T) {
	conf, err := sqlx.Connect("postgres", "host=localhost port=32798 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tr := repository.NewPgTodosRepository(conf)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		tr      domain.TodosRepository
		args    args
		want    *domain.Todo
		wantErr bool
	}{
		{"happy todo spec", tr, args{id: 0}, &domain.Todo{ID: 0, Done: false, Name: "ori", Details: "mehhh"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.tr.GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodoRepository.GetTodo() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgTodoRepository.GetTodo() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}
