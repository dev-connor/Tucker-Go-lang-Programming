package main

import (
	"fmt"
)

func divide(a, b int) {
	if b == 0 {
		panic("b 는 0 일 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
