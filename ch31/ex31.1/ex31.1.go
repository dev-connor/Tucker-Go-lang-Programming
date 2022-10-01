package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

type Success struct {
	Success bool `json:"success"`
}

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

type Todos []Todo

func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}

func (t Todos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux2.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

func GetTodoListHandler(response http.ResponseWriter, request *http.Request) {
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(response, http.StatusOK, list) // 2. ID 로 정렬하여 전체 목록 반환
}

func PostTodoHandler(response http.ResponseWriter, request *http.Request) {
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++ // 1. 새로운 ID 로 등록하고 만든
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(response, http.StatusCreated, todo)
}

func UpdateTodoHandler(response http.ResponseWriter, request *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(request.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(response, http.StatusOK, Success{true})
	} else {
		rd.JSON(response, http.StatusBadRequest, Success{false})
	}
}

func RemoveTodoHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(response, http.StatusOK, Success{true})
	} else {
		rd.JSON(response, http.StatusNotFound, Success{false})
	}
}

func main() {
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
