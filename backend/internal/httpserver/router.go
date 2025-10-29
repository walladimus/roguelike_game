package httpserver

import (
	"encoding/json"
	"net/http"

	"roguelike_game/backend/internal/httpserver/middleware"
	"roguelike_game/backend/internal/ws"

	"github.com/gorilla/mux"
)

const version = "0.0.1-dev"

func NewRouter() http.Handler {
	r := mux.NewRouter()

	// JSON health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"ok": true,
		})
	}).Methods(http.MethodGet)

	// Version endpoint
	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"version": version,
		})
	}).Methods(http.MethodGet)

	// Websocket test endpoint
	r.HandleFunc("/ws/echo", ws.EchoHandler)

	// Allow CORS
	return middleware.CORS(r)
}
 func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
 }