package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println("계산 중 입니다.")
	time.Sleep(3 * time.Second)
	fmt.Printf("%d 부터 %d 까지 합계는 %d 입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	wg.Wait()
	println("END")
}
