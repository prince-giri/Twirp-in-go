package main

import (
	"context"
	"log"
	"net/http"
	"sync"

	todo "github.com/prince-giri/twirp-todo/proto_generated"
	"github.com/twitchtv/twirp"
)

type TodoServer struct {
	mu    sync.Mutex
	todos []*todo.Todo
	idSeq int64
}

// CreateTodo creates a new todo item
func (s *TodoServer) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	log.Printf("⚡ CreateTodo called: %+v", req)
	if req.Title == "" || req.Description == "" {
		return nil, twirp.RequiredArgumentError("title or description missing")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.idSeq++
	t := &todo.Todo{
		Id:          s.idSeq,
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}
	s.todos = append(s.todos, t)

	return &todo.CreateTodoResponse{Todo: t}, nil
}

// GetTodo retrieves a todo by ID
func (s *TodoServer) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.GetTodoResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, t := range s.todos {
		if t.Id == req.Id {
			return &todo.GetTodoResponse{Todo: t}, nil
		}
	}
	return nil, twirp.NotFoundError("todo not found")
}

// ListTodos returns all todos
func (s *TodoServer) ListTodos(ctx context.Context, req *todo.ListTodosRequest) (*todo.ListTodosResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &todo.ListTodosResponse{Todos: s.todos}, nil
}

func main() {
	server := &TodoServer{}
	twirpHandler := todo.NewTodoServiceServer(server)

	log.Println("🚀 Twirp Todo server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", twirpHandler))
}
