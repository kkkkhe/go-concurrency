package main

import (
	"errors"
)

type Data struct {
	id   int
	name string
}

type Response struct {
	res Data
	err error
}

func HandleUrl(url string) (Data, error) {

	var d = map[string]Data{
		"first_url":  {1, "adsf"},
		"second_url": {2, "kkkkkkkk"},
	}
	data, has := d[url]
	if has {
		return data, nil
	}

	return Data{}, errors.New("There is no such url")
}

func Request(urls []string, done <-chan interface{}) <-chan Response {
	responses := make(chan Response)

	go func() {
		defer close(responses)
		for _, url := range urls {
			data, err := HandleUrl(url)
			res := Response{data, err}
			select {
			case <-done:
				return
			case responses <- res:
			}
		}
	}()

	return responses
}
