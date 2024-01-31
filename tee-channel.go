package main

func TeeChannel(done <-chan interface{}, stream <-chan int) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for value := range OrDone(done, stream) {
			for i := 0; i < 2; i++ {
				select {
				case out1 <- value:
				case out2 <- value:
				}
			}
		}
	}()

	return out1, out2
}
