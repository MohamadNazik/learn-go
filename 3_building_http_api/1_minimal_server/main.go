package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register the handler for the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on localhost:8080")

	// start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

// handler for the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
