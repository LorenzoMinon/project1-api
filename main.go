package main

import (
	"net/http"

	"github.com/LorenzoMinon/proyecto1-api/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.GetProducts)
	http.ListenAndServe(":8000", nil)
}
