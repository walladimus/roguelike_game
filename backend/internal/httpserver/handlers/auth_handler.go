package handlers

import (
	"encoding/json"
	"net/http"

	"roguelike_game/backend/internal/auth"
)

// AuthHandler provides minimal implementations for register/login (in-memory).
type AuthHandler struct {
	Service auth.Service
}

// Register handles POST /auth/register with minimal validation.
func (h AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req auth.Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.Service.Validate(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Login handles POST /auth/login with minimal validation.
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req auth.Credentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.Service.Validate(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"token": token})
}
