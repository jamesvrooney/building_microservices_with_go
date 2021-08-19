package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello handler for requests to /hello
type Hello struct {
	log *log.Logger
}

// NewHello creates a new Hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// SayHello say hello
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.log.Println("")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)

		return
	}

	fmt.Fprintf(w, "<h1>Hello James %s</h1>", d)
}
