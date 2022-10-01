package main

import (
	"encoding/json"
	mux2 "github.com/gorilla/mux"
	"net/http"
	"sort"
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

	// 2. 새 핸들러를 등록

	// 3. 임시 데이터 생성
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2

	return mux
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
