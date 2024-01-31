package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	fn := func() int { return 2 }

	res := Take(Repeat(done, fn), done, 10)

	for num := range res {
		if num == 2 {
			close(done)
		}
		fmt.Println(num)
	}

	time.Sleep(10 * time.Second)
}
