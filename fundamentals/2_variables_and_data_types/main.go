package main

import "fmt"

func main() {
	var name string = "John Doe"
	name2 := "Jane Doe" // only works inside functions
	var age int         // it becomes 0 by default

	// | Type    | Zero Value |
	// | ------- | ---------- |
	// | int     | 0          |
	// | string  | ""         |
	// | bool    | false      |
	// | slice   | nil        |
	// | map     | nil        |
	// | pointer | nil        |

	// Data types in Go:
	// - int, int8, int16, int32, int64
	// - uint, uint8, uint16, uint32, uint64
	// - float32, float64
	// - bool
	// - string
	// - array
	// - slice
	// - map
	// - struct

	age_ := 25
	price := 19.99
	active := true

	fmt.Println("Name:", name)
	fmt.Println("Name2:", name2)
	fmt.Println("Age:", age)
	fmt.Println("Age_:", age_)
	fmt.Println("Price:", price)
	fmt.Println("Active:", active)

	// go does not convert types implicitly, you need to convert them explicitly
	var a int = 10
	var b int64 = 5
	fmt.Println(a + int(b)) // convert b to int
}