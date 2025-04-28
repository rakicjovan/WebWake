package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rakicjovan/WebWake/services"
)

type WakeRequest struct {
	Mac string `json:"mac"`
}

func WakeHandler(w http.ResponseWriter, r *http.Request) {
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

	err = services.SendWOL(req.Mac)
	if err != nil {
		http.Error(w, "Failed to send WOL packet: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"success"}`))
}
