package main

import "fmt"

func main() {
	age := 25

	// if-else statement
	if age < 18 {
		fmt.Println("You are a minor.")
	}else{
		fmt.Printf("You are an adult.")
	}

	// if-else if statement

	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are an adult.")
	}


	// switch statement
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday":
		fmt.Println("It's the end of the week.")
	default:
		fmt.Println("It's a regular day.")
	}

	switch {
	case age < 13:
		fmt.Println("You are a child.")
	case age < 18:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are an adult.")
	}

	// loops
	// go only has for loop, but it can be used in different ways

	// for loop
	for i:=0; i<5; i++ {
		fmt.Println(i)
	}

	// while style for loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}


	for _,value := range []string{"Go", "is", "awesome"} {
		fmt.Println(value)
	}
}