package main

import (
	"errors"
	"fmt"
)

func main() {
	var addResult int = add(10,20)
	fmt.Println(addResult)

	var divideResult, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide Result:", divideResult)
	}
}

func add(a int, b int) int {
	return a + b
}

// short form
func addShort(a, b int) int {
	return a + b
}

// Multiple return values
func divide(a, b int) (int, error){
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}