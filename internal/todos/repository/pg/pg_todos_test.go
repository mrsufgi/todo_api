package repository_test

// NOT a unit test suit, uses a real db. prevent it from running with unit.
// improve test cases to make it run with unit\live\clearing db (not seeding!)
// requires a more complex setup (migrations etc.)
// these tests mostly helped with initial implementation

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrsufgi/todo_api/internal/domain"
	repository "github.com/mrsufgi/todo_api/internal/todos/repository/pg"
	log "github.com/sirupsen/logrus"
)

// TODO: create function to create PG connection from env variables so it works with docker/local pg.
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

func Test_pgTodosRepository_ReadTodo(t *testing.T) {
	conf, err := sqlx.Connect("postgres", "host=localhost port=35432 user=postgres password=postgres dbname=demo sslmode=disable")
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
			got, err := tt.tr.ReadTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodoRepository.ReadTodo() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgTodoRepository.ReadTodo() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_pgTodosRepository_SearchTodos(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id int
	}
	conn, err := sqlx.Connect("postgres", "host=localhost port=35432 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Todo
		wantErr bool
	}{
		{"happy todo spec", fields{conn: conn}, args{}, []domain.Todo{{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgTodosRepository(
				tt.fields.conn,
			)
			got, err := tr.SearchTodos()
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodosRepository.SearchTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgTodosRepository.SearchTodos() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: id is serial, running in parallel provide unexpected result.
// either create id externally (and mock it) or add a read by id and compare result.
func Test_pgTodosRepository_CreateTodo(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		todo domain.Todo
	}
	conn, err := sqlx.Connect("postgres", "host=localhost port=35432 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"happy todo spec", fields{conn: conn}, args{domain.Todo{Name: "test", Details: "test"}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgTodosRepository(
				tt.fields.conn,
			)
			got, err := tr.CreateTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodosRepository.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgTodosRepository.CreateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgTodosRepository_UpdateTodo(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id   int
		todo domain.Todo
	}
	conn, err := sqlx.Connect("postgres", "host=localhost port=35432 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy todo spec", fields{conn: conn},
			args{id: 1, todo: domain.Todo{Name: "updated", Details: "updated", Done: true}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgTodosRepository(
				tt.fields.conn,
			)
			got, err := tr.UpdateTodo(tt.args.id, tt.args.todo)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodosRepository.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgTodosRepository.UpdateTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgTodosRepository_DeleteTodo(t *testing.T) {
	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		id int
	}
	conn, err := sqlx.Connect("postgres", "host=localhost port=35432 user=postgres password=postgres dbname=demo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{"happy todo spec", fields{conn: conn}, args{id: 1}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := repository.NewPgTodosRepository(
				tt.fields.conn,
			)
			got, err := tr.DeleteTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgTodosRepository.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pgTodosRepository.DeleteTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}
