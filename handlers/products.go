package handlers

import (
	"encoding/json"
	"io"
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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body) //reads body and saves it.
	if err != nil {
		http.Error(w, "data error (no body request)", http.StatusBadRequest)
		return
	}

	var newProduct Product
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	data, err := json.Marshal(newProduct)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	foundIndex := -1
	id_str := r.PathValue("id")
	product_id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	for i := range products {
		if products[i].ID == product_id {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		http.Error(w, "invalid index", http.StatusInternalServerError)
		return
	}
	products = append(products[:foundIndex], products[foundIndex+1:]...)
	if err != nil {
		http.Error(w, "internal error server parsing to json", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	foundIndex := -1
	var updatedProduct Product
	// concrete variable, not a pointer!!! Unmarshal needs
	// a real address using >>> &
	str_id := r.PathValue("id")
	product_id, err := strconv.Atoi(str_id)
	if err != nil {
		http.Error(w, "cannot convert to int", http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &updatedProduct)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	newStock := updatedProduct.Stock
	for i := range products {
		if products[i].ID == product_id {
			foundIndex = i
			products[i].Stock = newStock
			break
		}
	}
	if foundIndex == -1 {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	data, err := json.Marshal(products[foundIndex])
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
