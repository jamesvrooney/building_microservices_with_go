package handlers

import (
	"fmt"
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
	fmt.Fprint(w, "<h1>Goodbye James!</h1>")
}
