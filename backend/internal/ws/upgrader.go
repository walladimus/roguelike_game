package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// shared upgrader used by all ws handlers in this package.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins during dev. Tighten later.
		return true
	},
}
