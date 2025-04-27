package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	http.HandleFunc("/wake", wakeHandler)

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

type WakeRequest struct {
	Mac string `json:"mac"`
}

func wakeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WakeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Mac == "" {
		http.Error(w, "Missing 'mac' field", http.StatusBadRequest)
		return
	}

	err = sendWOL(req.Mac)
	if err != nil {
		http.Error(w, "Failed to send WOL packet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"success"}`))
}
