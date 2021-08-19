package handlers

import (
	"jamesvrooney/microservices/data"
	"log"
	"net/http"
)

type Product struct {
	logger *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (product *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	products.ToJSON(w)
}
