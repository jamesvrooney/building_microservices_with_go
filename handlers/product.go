package handlers

import (
	"log"
	"net/http"

	"jamesvrooney/microservices/data"
)

type Product struct {
	logger *log.Logger
}

func NewProductHandler(logger *log.Logger) *Product {
	return &Product{logger}
}

func (p *Product) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJson(wr)

	if err != nil {
		http.Error(wr, "Unable to encode json", http.StatusInternalServerError)
	}
}
