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
	logger := log.New(os.Stdin, "product-api ", log.LstdFlags)

	helloHandler := handlers.NewHello(logger)
	goodbyeHandler := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()

	sm.Handle("/hello", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{
		Addr:         ":9898",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	interruptChannel := make(chan os.Signal, 1)

	signal.Notify(interruptChannel, os.Interrupt)
	signal.Notify(interruptChannel, os.Kill)

	// Block until a signal is received.
	interruptSignal := <-interruptChannel
	log.Println("Got signal:", interruptSignal)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(ctx)
}
