package main

import "fmt"

func main() {
	a := 25
	fmt.Println(update(a))
	fmt.Println("a:", a) // Original is unchanged because we passed a copy of a to the function

	// &age gives the memory address of age
	// *int is a pointer to an int, it can hold the memory address of an int variable
	var age int = 30
	fmt.Println(updateWithPointer(&age)) // Pass the address
	fmt.Println("age:", age) // Original is updated

	understandPointers(&age)
}

func update(a int) int {
	a = 100
	return a
}

func updateWithPointer(a *int) int {
	*a = 100 // *a dereferences the pointer to get the value and updates it
	return *a
}

func understandPointers(a *int) {
	fmt.Println("Value of a:", *a)
	fmt.Println("Address of a:", a)

	b := &a
	fmt.Println("Memory address of the a's memory address:", b)
}