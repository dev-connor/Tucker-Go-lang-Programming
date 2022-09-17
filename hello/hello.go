package main

func main() {
	b := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range b {

		for _, v := range arr {
			print(v, " ")
		}
	}
}
