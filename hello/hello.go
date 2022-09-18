package main

type Data struct {
	value int
	data  [200]int
}

func main() {
	p1 := &Data{}
	p2 := new(Data)

	println(p1, p2)
}
