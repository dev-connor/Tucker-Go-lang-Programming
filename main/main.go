package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case t := <-tick:
			fmt.Println("Tick", t)
		case end := <-terminate:
			fmt.Println("Terminated!", end)
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
