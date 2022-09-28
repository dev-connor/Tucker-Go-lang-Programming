package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello world") // 웹 핸들러 등록
	})

	http.ListenAndServe(":3000", nil) // 웹 서버 시작
}
