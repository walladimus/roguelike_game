package ws

import "encoding/json"

//generic Websocket payload shape used across ws
type Message struct {
	Type string `json:"type"`
	Data interface{} `json:"data"`
}

//convienience helper that converts message data into the provided pointer (v), by marshals then unmarshals (data is decoded by the stdlib as map[string]interface{})
func (m Message) UnmarshalData(v interface{}) error {
	b, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

//payload for a JOIN message
type JoinData struct {
	Username string `json:"username"`
}

//simple message payload
type ChatData struct {
	Text string `json:"text"`
}

//broadcast to clients to show who's in a lobby
type LobbyState struct {
	LobbyCode string `json:"lobbyCode"`
	Players []string `json:"players"`
}

//well known message types (string constants) - use to avoid typos
const (
	MessageTypeJoin = "JOIN"
	MessageTypeLeave = "LEAVE"
	MessageTypeChat = "CHAT"
	MessageTypeLobbyState = "LOBBYSTATE"
)