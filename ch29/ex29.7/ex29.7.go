package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World") // 인스턴스에 핸들러 등록
	})

	err := http.ListenAndServeTLS(":3000", "localhost.crt", "localhost.key", nil)

	// 1. HTTPS 서버 시작
	log.Fatal(err)
}
