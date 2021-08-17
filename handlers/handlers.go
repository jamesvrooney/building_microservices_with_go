package handlers

import (
	"fmt"
	"net/http"
)

// SayHello say hello
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello James...</h1>")
}
