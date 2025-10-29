package ws

import (
	"log"
	"net/http"
)

//echohandler handles GET /ws/echo
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	// HTTP -> Websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	for {
		//mt = message type
		//msg = data
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		log.Printf("ws recv: %s", string(msg))

		//echo back
		if err := conn.WriteMessage(mt, msg); err != nil {
			log.Printf("write error: %v", err)
			break
		}
	}
}