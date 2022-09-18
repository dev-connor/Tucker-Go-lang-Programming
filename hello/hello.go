package main

import "fmt"

type User struct {
	A int8
	C int8
	E int8
	B int
	D int
}

func main() {
	a := 10
	b := 20

	p1 := &a
	p2 := &a
	p3 := &b

	fmt.Printf("p1 == p2 : %t\n", p1 == p2)
	fmt.Printf("p2 == p3 : %t\n", p2 == p3)
}
