package main

import "fmt"

func InitPipeline(done chan interface{}, intValues ...int) <-chan int {
	resultSignal := make(chan int)

	go func() {
		defer close(resultSignal)
		for _, v := range intValues {
			select {
			case resultSignal <- v:
			case <-done:
				fmt.Println("return from init")
				return
			}
		}
	}()

	return resultSignal
}

func Multiply(values <-chan int, multiplier int, done <-chan interface{}) <-chan int {
	resultSignal := make(chan int)
	go func() {
		defer close(resultSignal)

		for value := range values {
			select {
			case <-done:
				return
			case resultSignal <- value * multiplier:
			}
		}
	}()
	return resultSignal
}

func Add(values <-chan int, additive int, done <-chan interface{}) <-chan int {
	resultSignal := make(chan int)

	go func() {
		defer close(resultSignal)
		for value := range values {
			select {
			case <-done:
				return
			case resultSignal <- value + additive:
			}
		}
	}()

	return resultSignal
}

//usage resultSignal := Add(Multiply(values, 2, done), 4, done)
