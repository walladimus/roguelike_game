package httpserver

import (
	"encoding/json"
	"net/http"

	"roguelike_game/backend/internal/httpserver/middleware"
	"roguelike_game/backend/internal/ws"
	"roguelike_game/backend/internal/storage"
	"github.com/gorilla/mux"
)

const version = "0.0.1-dev"

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_= json.NewEncoder(w).Encode(v)
}

func NewRouter(store storage.Store) http.Handler {
	r := mux.NewRouter()

	// JSON health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	}).Methods(http.MethodGet)

	// debug lobbies
	r.HandleFunc("/debug/lobbies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ws.SnapshotLobbies())
	}).Methods(http.MethodGet)

	//ws/lobby ? lobby=code
	r.HandleFunc("ws/lobby", ws.LobbyWSHandler)

	// Save/Resume
	r.HandleFunc("/api/save_resume", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var raw json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest); return
		}
		id, err := store.SaveResume(r.Context(), raw)
		if err != nil {
			http.Error(w, "save failed", http.StatusInternalServerError); return
		}
		writeJSON(w, http.StatusOK, map[string]string{"id": id})
	}).Methods(http.MethodPost)

	// /api/load_resume ? id= raw JSON blob
	r.HandleFunc("/api/load_resume", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" { http.Error(w, "missing id", http.StatusBadRequest); return }
		data, err := store.LoadResume(r.Context(), id)
		if err != nil { http.Error(w, "not found", http.StatusNotFound); return }
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}).Methods(http.MethodGet)

	return middleware.CORS(r)
}
