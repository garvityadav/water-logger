package handlers

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
