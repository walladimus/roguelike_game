package ws

import (
	"log"

	"roguelike_game/backend/internal/storage"
)

// helper that wraps a payload in the standard message envelope & sends to all clients
func (h *Hub) BroadcastJSON(t MessageType, payload interface{}) {
	data, err := MarshalMessage(t, payload)
	if err != nil {
		log.Printf("ws: failed to marshal message (%s): %v", t, err)
		return
	}
	h.Broadcast(data)
}

// BroadcastLobbyStateForCode sends a lobby snapshot to connected clients if the lobby hub exists.
func BroadcastLobbyStateForCode(lobbyCode string) {
	h, ok := lobbyManager.Get(lobbyCode)
	if !ok {
		return
	}
	h.BroadcastLobbyState(lobbyCode)
}

// BroadcastLobbyMembers sends a LOBBYSTATE snapshot based on lobby membership from storage.
// This keeps WS state aligned with HTTP lobby create/join/ready actions.
func BroadcastLobbyMembers(lobby storage.Lobby) {
	h, ok := lobbyManager.Get(lobby.Code)
	if !ok {
		return
	}

	memberStates := make([]LobbyMemberState, 0, len(lobby.Members))
	for _, m := range lobby.Members {
		name := m.Username
		if name == "" {
			name = "anonymous"
		}
		memberStates = append(memberStates, LobbyMemberState{
			Username: name,
			Ready:    m.Ready,
		})
	}

	h.BroadcastMessage(Message{
		Type: MessageTypeLobbyState,
		Data: LobbyState{
			LobbyCode: lobby.Code,
			Members:   memberStates,
		},
	})
}
