package main

import "fmt"

func main() {
	u := User{Name: "Alice"}
	introduce(u)
}

// Interfaces = a set of method signatures that a type can implement
type Speaker interface {
	Speak() string
}

// User implements the Speaker interface by defining the Speak method
type User struct {
	Name string
}

// Speak method for User type
func (u User) Speak() string {
	return "Hello, my name is " + u.Name
}

// introduce function takes a Speaker interface and calls its Speak method
func introduce(s Speaker) {
	fmt.Println(s.Speak())
}
