package main

import (
	"net/http"

	"github.com/LorenzoMinon/proyecto1-api/handlers"
)

func main() {
	http.HandleFunc("GET /products", handlers.GetProducts)
	http.HandleFunc("GET /products/{id}", handlers.GetProductByID)
	http.HandleFunc("POST /products", handlers.CreateProduct)
	http.HandleFunc("DELETE /products/{id}", handlers.DeleteProduct)
	http.ListenAndServe(":8000", nil)
}
