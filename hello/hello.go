package main

var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		println(m, s, g)
	}

	//m = s + 20 -> Error
}
