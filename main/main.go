package main

import "fmt"

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}

	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	result := func(a, b int) int {
		return a + b
	}(3, 4)

	fmt.Println(result)
}
