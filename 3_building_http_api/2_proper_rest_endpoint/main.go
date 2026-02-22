package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
	mu sync.Mutex
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", usersHandler)

	loggedMux := loggingMiddleware(mux)

	// log.Println("Server started on http://localhost:8080")
	// err := http.ListenAndServe(":8080", loggedMux)
	// if err != nil {
	// 	log.Fatal("Server failed to start:", err)
	// }

	// Graceful shutdown
	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	go func() {
		log.Println("Server started on http://localhost:8080")
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getUsers retrieves the list of users and sends it as a JSON response
func getUsers(w http.ResponseWriter) {

	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	mu.Lock()
	users = append(users, newUser)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		log.Println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
