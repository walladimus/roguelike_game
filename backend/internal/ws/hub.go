package ws

import (
	"sync"
	"encoding/json"
)

// Hub will manage all active websocket clients and broadcasts.
// Right now it's just a skeleton so the project builds.
type Hub struct {
	mu      sync.Mutex      // protects the below maps
	clients map[*Client]bool
}

// NewHub creates a new Hub with no clients yet.
func NewHub() *Hub {
	return &Hub{
		clients: make(map[*Client]bool),
	}
}

// Register adds a client to the hub.
// (We'll call this when someone connects to a lobby later.)
func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c] = true
}

// Unregister removes a client from the hub.
// (We'll call this when someone disconnects.)
func (h *Hub) Unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, c)
}

// Broadcast sends the given message to all registered clients.
// If a client's send buffer is full, we drop the client to avoid blocking the hub.
func (h *Hub) Broadcast(msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for c := range h.clients {
		select {
		case c.send <- msg:
			// queued successfully
		default:
			// client send buffer full or blocked — remove it to avoid blocking hub
			close(c.send)
			delete(h.clients, c)
		}
	}
}

// BroadcastMessage marshals a Message and sends it to all clients via Broadcast.
func (h *Hub) BroadcastMessage(m Message) {
	b, err := json.Marshal(m)
	if err != nil {
		// If marshaling fails, just drop it (shouldn't happen with our simple types).
		return
	}
	h.Broadcast(b)
}

// BroadcastLobbyState builds a LobbyState snapshot and broadcasts it.
// lobbyCode is passed so clients know which lobby this state describes.
func (h *Hub) BroadcastLobbyState(lobbyCode string) {
	// Collect player names under lock.
	h.mu.Lock()
	players := make([]string, 0, len(h.clients))
	for c := range h.clients {
		if c.Username != "" {
			players = append(players, c.Username)
		} else {
			players = append(players, "anonymous")
		}
	}
	h.mu.Unlock()

	h.BroadcastMessage(Message{
		Type: MessageTypeLobbyState,
		Data: LobbyState{
			LobbyCode: lobbyCode,
			Players:   players,
		},
	})
}
