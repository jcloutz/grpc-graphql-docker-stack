package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

const database string = "todo-app"
const collection string = "todos"

type Todo struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Complete  bool          `json:"complete" bson:"complete"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	h := &Handlers{
		db: session,
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/todos", h.GetAll).Methods("GET")
	router.HandleFunc("/api/todos/{id}", h.GetOne).Methods("GET")
	router.HandleFunc("/api/todos/{id}", h.UpdateOne).Methods("PUT")
	router.HandleFunc("/api/todos", h.CreateTodo).Methods("POST")

	http.ListenAndServe("localhost:8080", router)
}

type Handlers struct {
	db *mgo.Session
}

func (h *Handlers) GetAll(w http.ResponseWriter, r *http.Request) {
	sess := h.db.Copy()
	defer sess.Close()

	var results []Todo

	if err := sess.DB(database).C(collection).Find(nil).All(&results); err != nil {
		fmt.Println(err)
	}

	if results == nil {
		results = []Todo{}
	}

	h.Respond(w, results, 200)
}

func (h *Handlers) GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sess := h.db.Copy()
	defer sess.Close()

	var result Todo
	id := bson.ObjectIdHex(params["id"])
	if err := sess.DB(database).C(collection).FindId(id).One(&result); err != nil {
		h.RespondError(w, errors.New("Todo Not Found"), 404)

		return
	}

	h.Respond(w, result, 200)
}

func (h *Handlers) UpdateOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// find original todo
	var todo Todo
	id := bson.ObjectIdHex(params["id"])

	sess := h.db.Copy()
	defer sess.Close()

	if err := sess.DB(database).C(collection).FindId(id).One(&todo); err != nil {
		h.RespondError(w, errors.New("Todo Not Found"), 404)

		return
	}

	// Update todo
	decoder := json.NewDecoder(r.Body)
	var todoRequest Todo
	if err := decoder.Decode(&todoRequest); err != nil {
		panic(err)
	}

	todo.Name = todoRequest.Name
	todo.Complete = todoRequest.Complete
	todo.UpdatedAt = time.Now()

	err := sess.DB(database).C(collection).UpdateId(todo.ID, &todo)
	if err != nil {
		h.RespondError(w, errors.New("Unable to update todo"), 500)

		return
	}

	h.Respond(w, todo, 200)
}

type TodoCreateRequest struct {
	Name string `json:"name"`
}

func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todoRequest TodoCreateRequest
	if err := decoder.Decode(&todoRequest); err != nil {
		panic(err)
	}

	todo := Todo{
		ID:        bson.NewObjectId(),
		Name:      todoRequest.Name,
		Complete:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sess := h.db.Copy()
	defer sess.Close()

	sess.DB(database).C(collection).Insert(&todo)

	h.Respond(w, todo, 201)
}

// RespondError handle all error responses to the client
func (h *Handlers) RespondError(w http.ResponseWriter, apiError error, statusCode int) {

	e := ErrorResponse{
		Error: apiError.Error(),
	}

	h.Respond(w, e, statusCode)
}

// Respond handles all responses to the client
func (h *Handlers) Respond(w http.ResponseWriter, value interface{}, statusCode int) {

	js, err := json.Marshal(value)
	if err != nil {
		fmt.Println("error", err)
		js = []byte("{}")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	io.WriteString(w, string(js))
}
