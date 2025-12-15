package notices

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

// Store persists notices and player submitted requests.
type Store interface {
	ListNotices(ctx context.Context) ([]Notice, error)
	GetNotice(ctx context.Context, id string) (Notice, error)
	SaveRequest(ctx context.Context, input PlayerRequestInput) (PlayerRequest, error)
	ListRequests(ctx context.Context) ([]PlayerRequest, error)
}

type memoryStore struct {
	mu       sync.RWMutex
	notices  []Notice
	requests []PlayerRequest
}

// NewMemoryStore returns a simple in-memory store seeded with docs/changelog entries.
func NewMemoryStore() Store {
	return &memoryStore{
		notices:  defaultNotices(),
		requests: make([]PlayerRequest, 0),
	}
}

func (m *memoryStore) ListNotices(_ context.Context) ([]Notice, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return cloneNotices(m.notices), nil
}

func (m *memoryStore) GetNotice(_ context.Context, id string) (Notice, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, notice := range m.notices {
		if notice.ID == id {
			return notice, nil
		}
	}
	return Notice{}, ErrNoticeNotFound
}

func (m *memoryStore) SaveRequest(_ context.Context, input PlayerRequestInput) (PlayerRequest, error) {
	req := PlayerRequest{
		ID:        randomID(),
		Username:  input.Username,
		Message:   input.Message,
		CreatedAt: time.Now().UTC(),
	}

	m.mu.Lock()
	m.requests = append(m.requests, req)
	m.mu.Unlock()
	return req, nil
}

func (m *memoryStore) ListRequests(_ context.Context) ([]PlayerRequest, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]PlayerRequest, len(m.requests))
	copy(out, m.requests)
	return out, nil
}

func cloneNotices(in []Notice) []Notice {
	out := make([]Notice, len(in))
	copy(out, in)
	return out
}

func randomID() string {
	var buf [12]byte
	_, _ = rand.Read(buf[:])
	return base64.RawURLEncoding.EncodeToString(buf[:])
}

func defaultNotices() []Notice {
	return []Notice{
		{
			ID:          "backend-hardening",
			Title:       "Backend hardening + lobby sync",
			Summary:     "WebSocket upgrader enforces ALLOWED_ORIGINS and lobbies are now broadcast to connected clients automatically.",
			Body:        "From docs/changelog.md: \"WebSocket upgrader now checks ALLOWED_ORIGINS and logs connect/join/leave/state events. Lobby system extended with in-memory lobby storage (create/join/get/ready) and REST endpoints; HTTP router exposes /version and applies logging + recovery middleware.\"",
			Category:    "patch",
			Tags:        []string{"backend", "ws", "lobby"},
			PublishedAt: time.Date(2025, time.December, 8, 19, 55, 0, 0, time.UTC),
		},
		{
			ID:          "auth-updates",
			Title:       "Auth placeholders with hashed passwords",
			Summary:     "Register/login endpoints accept username+password, hash credentials, and fall back to in-memory storage unless ROGUE_DB_DSN is provided.",
			Body:        "From docs/changelog.md: \"Auth endpoints now hash passwords; use in-memory storage by default or Postgres when ROGUE_DB_DSN is set (requires migration 0002_add_user_password.sql); tokens are opaque, not JWT yet.\"",
			Category:    "patch",
			Tags:        []string{"auth", "storage"},
			PublishedAt: time.Date(2025, time.December, 8, 19, 55, 0, 0, time.UTC),
		},
		{
			ID:          "roadmap-update",
			Title:       "Roadmap + Play menu alignment",
			Summary:     "Phase 1 docs highlight the multi-lobby WebSocket endpoint (`/ws/lobby?lobby=CODE`) alongside the Play menu buttons (Play / Achievements / Friends / Settings / Notices / Coffee).",
			Body:        "From docs/roadmap.md and docs/changelog.md: \"Roadmap doc updated: Phase 1 WebSocket now described as multi-lobby endpoint with JOIN/LOBBYSTATE messaging; Phase 2 lobby channel notes revised; cleaned Play menu text.\"",
			Category:    "roadmap",
			Tags:        []string{"docs", "roadmap"},
			PublishedAt: time.Date(2025, time.December, 8, 19, 55, 0, 0, time.UTC),
		},
	}
}
