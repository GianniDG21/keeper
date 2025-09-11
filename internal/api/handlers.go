package api 

import (
	"encoding/json"
	"net/http"
)

//Function to check the server status
func (h *APIServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "available"})
}