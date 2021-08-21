package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Hello struct {
	logger *log.Logger
}

func NewHelloHandler(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello James Rooney! %s", time.Now())
}
