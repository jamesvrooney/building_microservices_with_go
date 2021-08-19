package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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
	encoder := json.NewEncoder(w)

	product := &Product{
		ID:      6,
		Name:    "Car (James Rooney)",
		Country: "Ireland",
	}

	encoder.Encode(product)
}

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}
