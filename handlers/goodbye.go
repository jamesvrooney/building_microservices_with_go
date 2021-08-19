package handlers

import (
	"log"
	"net/http"

	"jamesvrooney/microservices/data"
)

// Goodbye handler for requests to /hello
type Goodbye struct {
	log *log.Logger
}

// NewHello creates a new Hello handler
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// SayHello say hello
func (h *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	// encoder := json.NewEncoder(w)

	// encoder.Encode(products)
	products.ToJSON(w)
}
