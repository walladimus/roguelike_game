package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"roguelike_game/backend/internal/notices"
)

// NoticesHandler exposes the notices and player request endpoints.
type NoticesHandler struct {
	Service notices.Service
}

// ListNotices responds with all published notices.
func (h NoticesHandler) ListNotices(w http.ResponseWriter, r *http.Request) {
	notices, err := h.Service.ListNotices(r.Context())
	if err != nil {
		http.Error(w, "failed to load notices", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, notices)
}

// GetNotice responds with a single notice by id.
func (h NoticesHandler) GetNotice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	notice, err := h.Service.GetNotice(r.Context(), id)
	if err != nil {
		if errors.Is(err, notices.ErrNoticeNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "failed to load notice", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, notice)
}

// SubmitRequest accepts a player suggestion/bug report payload.
func (h NoticesHandler) SubmitRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var input notices.PlayerRequestInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	req, err := h.Service.SubmitRequest(r.Context(), input)
	if err != nil {
		if errors.Is(err, notices.ErrInvalidRequest) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "failed to save request", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusCreated, req)
}

// ListRequests returns every request submitted so far (admin view).
func (h NoticesHandler) ListRequests(w http.ResponseWriter, r *http.Request) {
	reqs, err := h.Service.ListRequests(r.Context())
	if err != nil {
		http.Error(w, "failed to load requests", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusOK, reqs)
}

func respondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
