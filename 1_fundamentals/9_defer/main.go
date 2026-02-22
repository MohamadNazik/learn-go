package main

import "fmt"

func main() {
	// defer is used to delay the execution of a function until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")
}
