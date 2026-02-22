package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // send value to channel
	}()

	value := <-ch // receive value from channel
	fmt.Println("Received:", value)

	// Using buffered channel (No need for a separate goroutine to send values)
	bufferedCh := make(chan string, 2) // Buffer size 2

	bufferedCh <- "Hello"
	bufferedCh <- "World"

	val1 := <-bufferedCh
	val2 := <-bufferedCh
	fmt.Println("Buffered Channel Values:", val1, val2)

	// Worker example
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	// send jobs and close the channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results) // Close results channel after all workers are done
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

	// Select example
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Pause execution for 2 seconds
		ch2 <- "Message from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	}

}

// <- chan receive only
// chan <- send only
func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
	}
}
