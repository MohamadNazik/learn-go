package main

import (
	"4_go_rest_api/db"
	"4_go_rest_api/handlers"
	"4_go_rest_api/middleware"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func loadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open .env file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}
}

func main() {
	loadEnv(".env")
	db.Connect()

	mux := http.NewServeMux()

	loggedMux := middleware.Logger(mux)

	// User routes
	mux.HandleFunc("GET /users", handlers.GetUsers)
	mux.HandleFunc("GET /users/{id}", handlers.GetUser)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	// Product routes
	mux.HandleFunc("GET /products", handlers.GetProducts)
	mux.HandleFunc("POST /products", handlers.CreateProduct)

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
