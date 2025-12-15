package storage

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"sync"
	"time"
)

var (
	ErrLobbyNotFound = errors.New("lobby not found")
	ErrUserExists    = errors.New("user already in lobby")
	ErrUserNotFound  = errors.New("user not in lobby")
	ErrMissingField  = errors.New("missing required field")
)

// blob that we can save/load later
type Resume struct {
	ID      string
	Data    json.RawMessage
	Updated time.Time
}

// lobby state used by HTTP + WS
type LobbyMember struct {
	Username string `json:"username"`
	Ready    bool   `json:"ready"`
}

type Lobby struct {
	Code    string        `json:"code"`
	Members []LobbyMember `json:"members"`
}

// minimal interface FE needs
type Store interface {
	SaveResume(ctx context.Context, data json.RawMessage) (string, error)
	LoadResume(ctx context.Context, id string) (json.RawMessage, error)
	CreateLobby(ctx context.Context, host string) (Lobby, error)
	JoinLobby(ctx context.Context, code, username string) (Lobby, error)
	GetLobby(ctx context.Context, code string) (Lobby, error)
	SetReady(ctx context.Context, code, username string, ready bool) (Lobby, error)
}

// in-memory implementation (no DB yet but coming later on)
type memStore struct {
	mu      sync.RWMutex
	data    map[string]Resume
	lobbies map[string]Lobby
}

// returns a fresh-store (for concurrent use)
func NewInMemory() Store {
	return &memStore{
		data:    make(map[string]Resume),
		lobbies: make(map[string]Lobby),
	}
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

func (m *memStore) CreateLobby(_ context.Context, host string) (Lobby, error) {
	if host == "" {
		return Lobby{}, ErrMissingField
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	code := genLobbyCode(m.lobbies)
	lobby := Lobby{
		Code:    code,
		Members: []LobbyMember{{Username: host, Ready: false}},
	}
	m.lobbies[code] = lobby
	return cloneLobby(lobby), nil
}

func (m *memStore) JoinLobby(_ context.Context, code, username string) (Lobby, error) {
	if code == "" || username == "" {
		return Lobby{}, ErrMissingField
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	lobby, ok := m.lobbies[code]
	if !ok {
		return Lobby{}, ErrLobbyNotFound
	}
	for _, mem := range lobby.Members {
		if mem.Username == username {
			return Lobby{}, ErrUserExists
		}
	}
	lobby.Members = append(lobby.Members, LobbyMember{Username: username, Ready: false})
	m.lobbies[code] = lobby
	return cloneLobby(lobby), nil
}

func (m *memStore) GetLobby(_ context.Context, code string) (Lobby, error) {
	if code == "" {
		return Lobby{}, ErrMissingField
	}

	m.mu.RLock()
	lobby, ok := m.lobbies[code]
	m.mu.RUnlock()
	if !ok {
		return Lobby{}, ErrLobbyNotFound
	}
	return cloneLobby(lobby), nil
}

func (m *memStore) SetReady(_ context.Context, code, username string, ready bool) (Lobby, error) {
	if code == "" || username == "" {
		return Lobby{}, ErrMissingField
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	lobby, ok := m.lobbies[code]
	if !ok {
		return Lobby{}, ErrLobbyNotFound
	}
	updated := false
	for i := range lobby.Members {
		if lobby.Members[i].Username == username {
			lobby.Members[i].Ready = ready
			updated = true
			break
		}
	}
	if !updated {
		return Lobby{}, ErrUserNotFound
	}
	m.lobbies[code] = lobby
	return cloneLobby(lobby), nil
}

func genID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	s := base64.RawURLEncoding.EncodeToString(b[:])
	return s
}

func genLobbyCode(existing map[string]Lobby) string {
	const letters = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	for {
		var code [6]byte
		for i := 0; i < 6; i++ {
			n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err != nil {
				// fallback to time-based seedless pseudo-random char
				code[i] = letters[int(time.Now().UnixNano()+int64(i))%len(letters)]
				continue
			}
			code[i] = letters[n.Int64()]
		}
		out := string(code[:])
		if _, ok := existing[out]; !ok {
			return out
		}
	}
}

func cloneLobby(l Lobby) Lobby {
	cpy := Lobby{Code: l.Code, Members: make([]LobbyMember, len(l.Members))}
	copy(cpy.Members, l.Members)
	return cpy
}
