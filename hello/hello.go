package main

import "fmt"

func main() {
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %t\n", x, x, x < x+1)
}
