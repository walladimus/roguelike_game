package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"roguelike_game/backend/internal/auth"
	"roguelike_game/backend/internal/httpserver/handlers"
	"roguelike_game/backend/internal/httpserver/middleware"
	"roguelike_game/backend/internal/notices"
	"roguelike_game/backend/internal/storage"
	"roguelike_game/backend/internal/ws"
)

const version = "0.0.1-dev"

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeLobbyError(w http.ResponseWriter, err error) {
	switch err {
	case storage.ErrMissingField:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case storage.ErrLobbyNotFound, storage.ErrUserNotFound:
		http.Error(w, err.Error(), http.StatusNotFound)
	case storage.ErrUserExists:
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func NewRouter(store storage.Store) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.RequestLogger, middleware.Recoverer)

	// services
	authStore := auth.UserStore(auth.NewInMemoryUserStore())
	if db, err := storage.NewPostgresDBFromEnv(); err == nil {
		authStore = auth.NewPostgresUserStore(db)
		log.Printf("auth: using postgres store")
	} else {
		log.Printf("auth: using in-memory store (%v)", err)
	}
	authService := auth.NewService(authStore)
	authHandler := handlers.AuthHandler{Service: authService}
	noticeStore := notices.NewMemoryStore()
	noticeService := notices.NewService(noticeStore)
	noticeHandler := handlers.NoticesHandler{Service: noticeService}

	// JSON health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"ok": true})
	}).Methods(http.MethodGet)

	// simple version endpoint to track build
	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"version": version})
	}).Methods(http.MethodGet)

	// debug lobbies
	r.HandleFunc("/debug/lobbies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ws.SnapshotLobbies())
	}).Methods(http.MethodGet)

	//ws/lobby ? lobby=code
	r.HandleFunc("/ws/lobby", ws.LobbyWSHandler).Methods(http.MethodGet)

	// Auth (stubbed for now)
	r.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	r.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)

	// Notices & player requests
	r.HandleFunc("/notices", noticeHandler.ListNotices).Methods(http.MethodGet)
	r.HandleFunc("/notices/{id}", noticeHandler.GetNotice).Methods(http.MethodGet)
	r.HandleFunc("/notices/requests", noticeHandler.ListRequests).Methods(http.MethodGet)
	r.HandleFunc("/notices/requests", noticeHandler.SubmitRequest).Methods(http.MethodPost)

	// Save/Resume
	r.HandleFunc("/api/save_resume", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var raw json.RawMessage
		if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		id, err := store.SaveResume(r.Context(), raw)
		if err != nil {
			http.Error(w, "save failed", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"id": id})
	}).Methods(http.MethodPost)

	// /api/load_resume ? id= raw JSON blob
	r.HandleFunc("/api/load_resume", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		data, err := store.LoadResume(r.Context(), id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}).Methods(http.MethodGet)

	// Lobby create
	r.HandleFunc("/lobby/create", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body struct {
			Username string `json:"username"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		lobby, err := store.CreateLobby(r.Context(), body.Username)
		if err != nil {
			writeLobbyError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, lobby)
		ws.BroadcastLobbyMembers(lobby)
	}).Methods(http.MethodPost)

	// Lobby join
	r.HandleFunc("/lobby/join", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body struct {
			Code     string `json:"code"`
			Username string `json:"username"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		lobby, err := store.JoinLobby(r.Context(), body.Code, body.Username)
		if err != nil {
			writeLobbyError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, lobby)
		ws.BroadcastLobbyMembers(lobby)
	}).Methods(http.MethodPost)

	// Lobby ready toggle
	r.HandleFunc("/lobby/{code}/ready", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body struct {
			Username string `json:"username"`
			Ready    bool   `json:"ready"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		code := mux.Vars(r)["code"]
		lobby, err := store.SetReady(r.Context(), code, body.Username, body.Ready)
		if err != nil {
			writeLobbyError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, lobby)
		ws.BroadcastLobbyMembers(lobby)
	}).Methods(http.MethodPost)

	// Lobby detail
	r.HandleFunc("/lobby/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := mux.Vars(r)["code"]
		lobby, err := store.GetLobby(r.Context(), code)
		if err != nil {
			writeLobbyError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, lobby)
	}).Methods(http.MethodGet)

	return middleware.CORS(r)
}
