package handlers

import (
	"4_go_rest_api/db"
	"4_go_rest_api/models"
	"encoding/json"
	"net/http"
)

// GET /products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(`SELECT id, title, description, price, created_at from products`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	products := []models.Product{}

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO products (title, description, price) VALUES (?, ?, ?)`

	result, err := db.DB.Exec(query, p.Title, p.Description, p.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	p.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
