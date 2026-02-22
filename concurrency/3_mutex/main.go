package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

func increment() {
	defer wg.Done()

	mu.Lock() // prevents concurrent writes
	count++
	mu.Unlock()
}

func main() {
	for range 1000 {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Counter:", count)
}
