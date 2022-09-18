package main

type Data struct {
	value int
	data  [200]int
}

func main() {
	str := "Hello World"
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	println(str)
	println(string(runes))
}
