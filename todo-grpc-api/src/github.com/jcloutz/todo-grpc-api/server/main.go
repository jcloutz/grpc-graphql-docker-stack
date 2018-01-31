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
	port       = ":443"
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

func RequestStreamLogger() grpc.ServerOption {
	return grpc.StreamInterceptor(requestStreamLogger)
}

func requestStreamLogger(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()

	err := handler(srv, ss)
	log.Printf("invoke server method=%s duration=%s error=%s", info.FullMethod, time.Since(start), err)

	return err
}

func RequestLogger() grpc.ServerOption {
	return grpc.UnaryInterceptor(requestLogger)
}

func requestLogger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("invoke server method=%s duration=%s error=%s", info.FullMethod, time.Since(start), err)

	return resp, err
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

	// Set up TLS
	// cert, err := tls.LoadX509KeyPair("/run/secrets/grpc_api_crt", "/run/secrets/grpc_api_key")
	// if err != nil {
	// 	log.Fatalf("Error loading cert: %s", err)
	// 	return
	// }
	// config := tls.Config{
	// 	Certificates: []tls.Certificate{cert},
	// 	// InsecureSkipVerify: true,
	// }

	// serverOption := grpc.Creds(credentials.NewTLS(&config))
	// s := grpc.NewServer(serverOption, RequestLogger(), RequestStreamLogger())
	s := grpc.NewServer(RequestLogger(), RequestStreamLogger())
	todoPb.RegisterTodoServer(s, &service)
	log.Println("Listening on port 8081")
	s.Serve(lis)
}
