package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("output.txt") // 1. 파일 생성
	if err != nil {
		fmt.Errorf("Create: %v\n", err)
		return
	}

	defer f.Close()

	const name, age = "Kim", 22
	n, err := fmt.Fprint(f, name, " is ", age, " years old\n") // 2.

	if err != nil {
		fmt.Errorf("Fprint: %v\n", err)
	}
	print(n, " bytes written.\n")
}
