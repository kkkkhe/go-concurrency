package main

import (
	"fmt"
	"math/rand"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	fn := func() int { return rand.Intn(100) }

	res := Take(Repeat(done, fn), done, 11)

	out1, out2 := TeeChannel(done, res)

	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
