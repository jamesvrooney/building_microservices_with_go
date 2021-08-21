package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"jamesvrooney/microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "products-api ", log.LstdFlags)

	helloHandler := handlers.NewHelloHandler(logger)
	productHandler := handlers.NewProductHandler(logger)

	sm := http.NewServeMux()

	sm.Handle("/hello", helloHandler)
	sm.Handle("/product", productHandler)

	server := http.Server{
		Addr:         ":8989",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,  // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logger.Printf("Error starting the server %s\n", err)

			os.Exit(1)
		}
	}()

	interruptChannel := make(chan os.Signal)

	signal.Notify(interruptChannel, os.Interrupt)
	signal.Notify(interruptChannel, os.Kill)

	interruptSignal := <-interruptChannel

	logger.Println("Received interrupt signal", interruptSignal)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	server.Shutdown(ctx)
}
