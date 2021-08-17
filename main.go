package main

import (
	"net/http"

	"jamesvrooney/microservices/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SayHello)

	http.ListenAndServe(":9898", nil)
}
