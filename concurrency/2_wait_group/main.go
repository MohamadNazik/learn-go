package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // Expecting one goroutine

		// go functName(i)
		go func(i int) {
			defer wg.Done() // Signal that this goroutine is done
			fmt.Println(i)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to complete
}
