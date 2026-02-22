package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // Simulate slow work
		result <- "Hello, World!"
	}()

	select {
	case res := <-result:
		println(res)
	case <-ctx.Done():
		println("Operation timed out") // output: Operation timed out
	}
}
