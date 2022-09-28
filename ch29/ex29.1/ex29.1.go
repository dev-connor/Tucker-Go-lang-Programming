package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World") // 인스턴스에 핸들러 등록
	})
	mux.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Bar")
	})

	http.ListenAndServe(":3000", mux) // 웹 서버 시작
}
