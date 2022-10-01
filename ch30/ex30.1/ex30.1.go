package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"net/http"
	"sort"
	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student // 학생 목록을 저장하는 맵
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux2.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	// 2. 새 핸들러를 등록

	// 3. 임시 데이터 생성
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
}

func DeleteStudentHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	delete(students, id)
	response.WriteHeader(http.StatusOK)
}

func PostStudentHandler(response http.ResponseWriter, request *http.Request) {
	var student Student
	err := json.NewDecoder(request.Body).Decode(&student) // JSON 데이터 변환
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++ // 2. id 를 증가시킨 후 맵에 등록
	student.Id = lastId
	students[lastId] = student
	response.WriteHeader(http.StatusCreated)
}

func GetStudentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux2.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]

	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		// 2. id 에 해당하는 학생이 없으면 에러
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(student)
}

func GetStudentListHandler(writer http.ResponseWriter, request *http.Request) {
	list := make(Students, 0) // 1. 학생 목록을 Id 로 정렬
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(list) // 2. JSON 포맷으로 변경
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
