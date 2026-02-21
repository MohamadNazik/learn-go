package main

import "os"

func main() {
	file, err := os.Create("./fundamentals/10_file_handling/file.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hello")
}
