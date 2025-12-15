package ws

import "encoding/json"

// high-level type of WebSocket message.
type MessageType string

// well known message types (string constants) - use to avoid typos
const (
	// sent by client when they join a lobby.
	MessageTypeJoin MessageType = "JOIN"

	// sent by client or server to indicate leaving.
	MessageTypeLeave MessageType = "LEAVE"

	// chat message sent by a client, broadcast by the server.
	MessageTypeChat MessageType = "CHAT"

	// server -> clients: snapshot of the lobby
	MessageTypeLobbyState MessageType = "LOBBYSTATE"

	// server -> clients: snapshot of the current turn state for the lobby
	MessageTypeTurnState MessageType = "TURNSTATE"
)

// generic Websocket payload shape used across ws
type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

// convienience helper that converts message data into the provided pointer (v), by marshals then unmarshals (data is decoded by the stdlib as map[string]interface{})
func (m Message) UnmarshalData(v interface{}) error {
	b, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

// convienience helper that wraps a payload and type into a Message and marshals to JSON bytes
func MarshalMessage(t MessageType, payload interface{}) ([]byte, error) {
	msg := Message{
		Type: t,
		Data: payload,
	}
	return json.Marshal(msg)
}

// payload for a JOIN message
type JoinData struct {
	Username string `json:"username"`
}

// simple message payload
type ChatData struct {
	Text string `json:"text"`
}

// broadcast to clients to show who is in a lobby
type LobbyState struct {
	LobbyCode string             `json:"lobbyCode"`
	Members   []LobbyMemberState `json:"members"`
}

// LobbyMemberState mirrors ready/username for WS payloads.
type LobbyMemberState struct {
	Username string `json:"username"`
	Ready    bool   `json:"ready"`
}

// broadcast to clients to show the current turn state for a lobby
// (State can later be replaced with a concrete DTO instead of interface{}).
type TurnState struct {
	LobbyCode string      `json:"lobbyCode"`
	State     interface{} `json:"state"`
}
