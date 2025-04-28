package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/rakicjovan/WebWake/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/wake", handlers.WakeHandler)

	// Read PORT environment variable, fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := ":" + port
	log.Printf("Starting server on %s\n", address)

	// Setup CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := corsHandler.Handler(http.DefaultServeMux)

	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
