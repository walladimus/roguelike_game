package storage

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// blob that we can save/load later
type Resume struct {
	ID string
	Data json.RawMessage
	Updated time.Time
}

// minimal interface FE needs
type Store interface {
	SaveResume(ctx context.Context, data json.RawMessage) (string, error)
	LoadResume(ctx context.Context, id string) (json.RawMessage, error)
}

// in-memory implementation (no DB yet but coming later on)
type memStore struct {
	mu sync.RWMutex
	data map[string]Resume
}

// returns a fresh-store (for concurrent use)
func NewInMemory() Store {
	return &memStore{data: make(map[string]Resume)}
}

func (m *memStore) SaveResume(_ context.Context, data json.RawMessage) (string, error) {
	id := genID()
	m.mu.Lock()
	m.data[id] = Resume{ID: id, Data: data, Updated: time.Now().UTC()}
	m.mu.Unlock()
	return id, nil
}

func (m *memStore) LoadResume(_ context.Context, id string) (json.RawMessage, error) {
	m.mu.RLock()
	r, ok := m.data[id]
	m.mu.RUnlock()
	if !ok {
		return nil, errors.New("not found")
	}
	out := make([]byte, len(r.Data))
	copy(out, r.Data)
	return out, nil
}

func genID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	s := base64.RawURLEncoding.EncodeToString(b[:])
	return s
}
