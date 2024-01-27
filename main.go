package main

import "fmt"

func main() {

	done := make(chan interface{})
	defer close(done)

	urls := []string{"first_url", "second_ur"}
	resp := Request(urls, done)

	for res := range resp {
		if res.err != nil {
			fmt.Println("Error occured: ", res.err)
			continue
		}
		fmt.Println(res.res)
	}
}
