package main

import "fmt"

func main() {

	// Slices = pointers + length + capacity

	numbers := []int{1, 2, 3, 4, 5} // A slice of integers
	fmt.Println("Numbers:", numbers)

	var nums [3]int            // An array of 3 integers, it has a fixed size
	fmt.Println("Nums:", nums) // Nums: [0 0 0]
	nums[0] = 10
	fmt.Println("Nums after update:", nums) // Nums after update: [10 0 0]

	fmt.Println("numbers[1:3]:", numbers[1:3])

	numbers = append(numbers, 6, 7) // Append new elements to the slice
	fmt.Println("Numbers after append:", numbers)

	// Slices has length and capacity
	fmt.Println("Length of numbers:", len(numbers))
	fmt.Println("Capacity of numbers:", cap(numbers))

	// Maps
	// why make?
	// Because maps zero's value is nil, Nil map cannot store values
	ages := make(map[string]int) // make creates a new map (string keys and int values)

	ages["Alice"] = 30
	ages["Bob"] = 25

	value, ok := ages["John"] // ok is true if the key exists
	if ok {
		fmt.Println("John's age:", value)
	} else {
		fmt.Println("John's age not found")
	}

}
