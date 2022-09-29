package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(writer http.ResponseWriter, request *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)
	writer.Header().Add("content-type", "application/json") // 3. JSON 포맷임을 표시
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(data))
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
