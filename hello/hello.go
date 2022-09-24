package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	println(n)
}
