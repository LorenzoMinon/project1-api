package handlers

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

var nextID = 4

var products = []Product{
	{ID: 1, Name: "Jeans", Price: 100.00, Description: "Levis 505", Stock: 5},
	{ID: 2, Name: "Black Jeans", Price: 105.00, Description: "Levis 505 black", Stock: 4},
	{ID: 3, Name: "White Jeans", Price: 104.00, Description: "Levis 505 white", Stock: 2},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set header
	data, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
