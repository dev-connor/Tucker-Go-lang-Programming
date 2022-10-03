package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.IntSlice(s))

	fmt.Println(s)
}
