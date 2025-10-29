package ws

import (
	"net/http"
)

var lobbyManager = NewLobbyManager()
var globalHub = NewHub()

func LobbyWSHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("lobby")
	if code == "" {
		code = "default"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	//create client that self-registers and starts pumps
	hub := lobbyManager.GetOrCreate(code) 
	NewClient(hub, conn, code)
}