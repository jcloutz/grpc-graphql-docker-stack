package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"

	todoPb "github.com/jcloutz/todo-grpc-api/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	// Set up TLS
	cert, err := tls.LoadX509KeyPair("/Users/jeremycloutier/develop/tech-talks/docker-swarm/secrets/client.crt", "/Users/jeremycloutier/develop/tech-talks/docker-swarm/secrets/client.key")
	if err != nil {
		log.Fatalf("Error loading cert: %s", err)
		return
	}
	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	creds := credentials.NewTLS(&config)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	// conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := todoPb.NewTodoClient(conn)

	createTodo(client, &todoPb.TodoCreateRequest{
		Name:     "It works!",
		Complete: false,
	})

	// updateTodo(client, &todoPb.TodoUpdateRequest{
	// 	Id:       "5a70b1ac541f30094ce4542f",
	// 	Name:     "gRPC Works!",
	// 	Complete: true,
	// })

	getTodos(client, &todoPb.TodoFilter{})
}
