package main

import (
	"fmt"
)

func main() {

	// Run normally
	sayHello()

	// Run as a goroutine
	go sayHello() // go keyword makes it run concurrently

	// Race condition example
	for range 1000 {
		go increment() // multiple goroutines incrementing the same counter
	}

	fmt.Println("Counter:", counter) // May not print 1000 due to race condition
}

func sayHello() {
	fmt.Println("Hello, World!")
}

var counter int

func increment() {
	counter++
}
