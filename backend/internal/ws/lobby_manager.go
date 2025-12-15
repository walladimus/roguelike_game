package ws

import (
	"sync"
)

// manages multiple Hub instances keyed by lobby code
type LobbyManager struct {
	mu   sync.Mutex
	hubs map[string]*Hub
}

// constructs a manager
func NewLobbyManager() *LobbyManager {
	return &LobbyManager{
		hubs: make(map[string]*Hub),
	}
}

// returns the hub for a given code (created if missing)
func (m *LobbyManager) GetOrCreate(code string) *Hub {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.hubs[code]
	if ok {
		return h
	}

	h = NewHub()
	m.hubs[code] = h
	return h
}

// Get returns the hub for a given code if it exists.
func (m *LobbyManager) Get(code string) (*Hub, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.hubs[code]
	return h, ok
}

// removes the hub for code if no clients are active (returns true)
func (m *LobbyManager) RemoveIfEmpty(code string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	h, ok := m.hubs[code]
	if !ok {
		return false
	}

	//access h.client safely
	h.mu.Lock()
	empty := len(h.clients) == 0
	h.mu.Unlock()

	if empty {
		delete(m.hubs, code)
		return true
	}
	return false
}

// calls fn for every lobby under a lock
func (m *LobbyManager) ForEach(fn func(code string, h *Hub)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for code, h := range m.hubs {
		fn(code, h)
	}
}
