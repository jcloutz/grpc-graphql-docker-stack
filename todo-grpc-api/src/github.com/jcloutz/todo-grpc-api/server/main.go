package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"

	"github.com/golang/protobuf/ptypes"
	todoPb "github.com/jcloutz/todo-grpc-api/todo"
	"gopkg.in/mgo.v2/bson"
)

const (
	port       = ":8081"
	database   = "todo-app"
	collection = "todos"
)

type Todo struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Complete  bool          `json:"complete" bson:"complete"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

func todoToProtobuf(todo *Todo) *todoPb.TodoResponse {
	createdAt, _ := ptypes.TimestampProto(todo.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(todo.UpdatedAt)

	return &todoPb.TodoResponse{
		Id:        todo.ID.Hex(),
		Name:      todo.Name,
		Complete:  todo.Complete,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type todoService struct {
	db *mgo.Session
}

func (t *todoService) GetTodos(filter *todoPb.TodoFilter, stream todoPb.Todo_GetTodosServer) error {
	sess := t.db.Copy()
	defer sess.Close()

	var results []Todo
	if err := sess.DB(database).C(collection).Find(nil).All(&results); err != nil {
		fmt.Println(err)
	}

	if results == nil {
		results = []Todo{}
	}

	for _, todo := range results {
		if err := stream.Send(todoToProtobuf(&todo)); err != nil {
			return err
		}
	}

	return nil
}

func (t *todoService) CreateTodo(ctx context.Context, in *todoPb.TodoCreateRequest) (*todoPb.TodoResponse, error) {
	todo := Todo{
		ID:        bson.NewObjectId(),
		Name:      in.Name,
		Complete:  in.Complete,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}

	sess := t.db.Copy()
	defer sess.Close()

	if err := sess.DB(database).C(collection).Insert(&todo); err != nil {
		return nil, err
	}

	return todoToProtobuf(&todo), nil
}

func (t *todoService) UpdateTodo(ctx context.Context, in *todoPb.TodoUpdateRequest) (*todoPb.TodoResponse, error) {
	sess := t.db.Copy()
	defer sess.Close()

	var todo Todo
	id := bson.ObjectIdHex(in.Id)

	if err := sess.DB(database).C(collection).FindId(id).One(&todo); err != nil {
		return nil, err
	}

	todo.Name = in.Name
	todo.Complete = in.Complete
	todo.UpdatedAt = time.Now()

	if err := sess.DB(database).C(collection).UpdateId(todo.ID, &todo); err != nil {
		return nil, err
	}

	return todoToProtobuf(&todo), nil
}

func main() {
	session, err := mgo.Dial("database")
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := todoService{
		db: session,
	}
	s := grpc.NewServer()
	todoPb.RegisterTodoServer(s, &service)
	log.Println("Listening on port 8081")
	s.Serve(lis)
}
