package ws

import (
	"log"
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
		log.Printf("ws: upgrade failed lobby=%s err=%v", code, err)
		return
	}
	log.Printf("ws: connected lobby=%s remote=%s", code, r.RemoteAddr)

	//create client that self-registers and starts pumps
	hub := lobbyManager.GetOrCreate(code)
	NewClient(hub, conn, code)
}
