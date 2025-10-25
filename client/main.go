package main

import (
	"context"
	"log"
	"net/http"

	todo "github.com/prince-giri/twirp-todo/proto_generated"
)

func main() {
	// Use a real HTTP client
	client := todo.NewTodoServiceProtobufClient("http://localhost:8080", http.DefaultClient)

	resp, err := client.CreateTodo(context.Background(), &todo.CreateTodoRequest{
		Title:       "Learn Go",
		Description: "Study Twirp",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created todo:", resp.Todo)
}
