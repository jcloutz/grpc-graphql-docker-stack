package main

import (
	"context"
	"io"
	"log"

	todoPb "github.com/jcloutz/todo-grpc-api/todo"
	"google.golang.org/grpc"
)

const address = "localhost:8081"

func getTodos(client todoPb.TodoClient, filter *todoPb.TodoFilter) {
	stream, err := client.GetTodos(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get todos: %v", err)
	}

	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.GetTodos(_) = _, %v", client, err)
		}
		log.Printf("Todo: %v", todo)
	}
}

func createTodo(client todoPb.TodoClient, todo *todoPb.TodoCreateRequest) {
	resp, err := client.CreateTodo(context.Background(), todo)
	if err != nil {
		log.Fatalf("Could not create todo: %v", err)
	}

	log.Printf("Todo created: %v", resp)
}

func updateTodo(client todoPb.TodoClient, update *todoPb.TodoUpdateRequest) {
	resp, err := client.UpdateTodo(context.Background(), update)
	if err != nil {
		log.Fatalf("Could not update todo: %v", err)
	}

	log.Printf("Todo updated: %v", resp)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := todoPb.NewTodoClient(conn)

	updateTodo(client, &todoPb.TodoUpdateRequest{
		Id:       "5a70b1ac541f30094ce4542f",
		Name:     "gRPC Works!",
		Complete: true,
	})

	getTodos(client, &todoPb.TodoFilter{})
}
