package main

import (
	"sync"
)

func Take(numsStream <-chan int, done chan interface{}, amount int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for i := 0; i < amount; i++ {
			select {
			case <-done:
				return
			case result <- <-numsStream:
			}
		}
	}()
	return result
}

func PrimeFinder(done <-chan interface{}, intStream <-chan int) <-chan int {
	primeStream := make(chan int)
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer -= 1
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

func Repeat(done <-chan interface{}, fn func() int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case result <- fn():
			}
		}
	}()
	return result
}

func Fanin(done <-chan interface{}, streams ...<-chan int) <-chan int {
	resultStream := make(chan int)
	wg := sync.WaitGroup{}
	var mult = func(stream <-chan int) {
		defer wg.Done()
		for num := range stream {
			select {
			case <-done:
				return
			case resultStream <- num:
			}
		}
	}
	wg.Add(len(streams))
	for _, stream := range streams {
		go mult(stream)
	}
	go func() {
		wg.Wait()
		defer close(resultStream)
	}()
	return resultStream
}
