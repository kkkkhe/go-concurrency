package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	numStream := make(chan interface{}, 100)
	for i := 0; i < 100; i++ {
		numStream <- i
	}

	done := make(chan interface{})
	intRand := func() int { return rand.Intn(50000000) }
	now := time.Now()
	streams := make([]<-chan int, 8)

	for i := 0; i < 8; i++ {
		streams[i] = PrimeFinder(done, Repeat(done, intRand))
	}
	for r := range Take(Fanin(done, streams...), done, 10) {
		fmt.Println("Result num: ", r)
	}
	fmt.Println("Since:", time.Since(now).Seconds())
	close(done)
}
