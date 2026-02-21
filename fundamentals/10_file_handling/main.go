package main

import "os"

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hello")
}
