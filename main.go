package main

import "fmt"

func main() {

	done := make(chan interface{})

	values := InitPipeline(done, 1, 2, 3, 4, 5)
	result := Add(Multiply(values, 2, done), 4, done)
	// done <- true
	for value := range result {
		fmt.Println(value)
	}

}
