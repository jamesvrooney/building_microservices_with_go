package main

import (
	"log"
	"net/http"
	"os"

	"jamesvrooney/microservices/handlers"
)

func main() {
	logger := log.New(os.Stdin, "product-api ", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)

	sm := http.NewServeMux()

	sm.Handle("/hello", helloHandler)

	http.ListenAndServe(":9898", sm)
}
