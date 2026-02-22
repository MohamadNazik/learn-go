package handlers

import (
	"4_go_rest_api/db"
	"4_go_rest_api/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /users

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, email, created_at FROM users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var u models.User

	query := "SELECT id, name, email, created_at FROM users WHERE id = ?"
	err := db.DB.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)

}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	result, err := db.DB.Exec(query, u.Name, u.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	u.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// PUT /users/{id}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
	_, err := db.DB.Exec(query, u.Name, u.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id_int, _ := strconv.Atoi(id)
	u.ID = id_int
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	query := "DELETE FROM users WHERE id = ?"
	_, err := db.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
