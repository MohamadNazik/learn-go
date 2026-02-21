package main

import "fmt"

func main() {
	u := User{ID: 1, Name: "Alice"}
	fmt.Println("User Name:", u.Name)
}

// Struct = Custom data type
type User struct {
	ID   int
	Name string
}

// Attach behavior to struct
func (u User) Greet() string {
	return "Hello, " + u.Name
}
