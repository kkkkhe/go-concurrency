package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func printGreeting(ctx context.Context) error {
	greeting, err := genGreeting(ctx)
	if err != nil {
		return errors.New("context deadline exceeded")
	}
	fmt.Printf("%s world!\n", greeting)
	return nil
}
func printFarewell(ctx context.Context) error {
	farewell, err := genFarewell(ctx)
	if err != nil {
		return errors.New("context canceled")
	}
	fmt.Printf("%s world!\n", farewell)
	return nil

}
func genGreeting(ctx context.Context) (string, error) {
	ctx2, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	fmt.Println("here")
	time.Sleep(4 * time.Second)
	fmt.Println("her1")
	l, err := locale(ctx2)
	if err != nil {
		return "", err
	}
	if l == "ua" {
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(ctx context.Context) (string, error) {
	l, err := locale(ctx)
	if err != nil {
		return "", err
	}
	if l == "ua" {
		return "goodbye", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func locale(ctx context.Context) (string, error) {
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Println(deadline.Sub(time.Now()))
		if deadline.Sub(time.Now().Add(6*time.Second)) <= 0 {
			fmt.Println("here")
			return "", context.DeadlineExceeded
		}
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(1 * time.Minute):
	}
	return "ua", nil
}

/*
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(ctx); err != nil {
			cancel()
			fmt.Println("Cannot print greeting: ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(ctx); err != nil {
			fmt.Println("Cannot print farewell: ", err)
		}
	}()

	wg.Wait()
}
*/
