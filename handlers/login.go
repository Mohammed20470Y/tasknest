package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mohammed20470Y/tasknest/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler returns a JWT token on successful dummy login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hardcoded credentials for now
	if req.Username != "admin" || req.Password != "password" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
