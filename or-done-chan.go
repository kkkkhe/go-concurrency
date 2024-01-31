package main

func OrDone(done <-chan interface{}, stream <-chan int) <-chan interface{} {
	resultStream := make(chan interface{})

	go func() {
		defer close(resultStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-stream:
				if ok == false {
					return
				}
				select {
				case <-done:
				case resultStream <- v:
				}
			}
		}
	}()

	return resultStream
}

//usage

// stream := make(chan interface{}, 10)

// for i := 0; i < 10; i++ {
// 	stream <- i
// }
// done := make(chan interface{})
// result := OrDone(done, stream)

// for c := range result {
// 	if c == 5 {
// 		close(done)
// 	}
// 	fmt.Println(c)
// }
