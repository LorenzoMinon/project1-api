package main

import (
	"net/http"

	"github.com/LorenzoMinon/proyecto1-api/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.GetProducts)
	http.HandleFunc("/products/{id}", handlers.GetProductByID)
	http.ListenAndServe(":8000", nil)
}
