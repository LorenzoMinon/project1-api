package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	var foundProduct *Product
	id_str := r.PathValue("id")
	product_id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i := range products {
		if products[i].ID == product_id {
			foundProduct = &products[i]
		}
	}
	if foundProduct == nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	data, err := json.Marshal(foundProduct)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
