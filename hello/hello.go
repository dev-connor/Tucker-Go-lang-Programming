package main

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); {
	case age < 10:
		println("Child")
	case age < 20:
		println("Teenager")
	case age < 30:
		println("20s")
	default:
		println("My age is", age) // age 값 사용
	}
}
