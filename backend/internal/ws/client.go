package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	//writewait = time allowed to write a message
	//pongwait = time allowed to read the message
	//pingPeriod = time to send pings
	writeWait = 10 * time.Second
	pongWait = 60 *time.Second
	pingPeriod = (pongWait * 9)/10
	maxMessageSize = 512
)

//Client represents some connected websocket client
type Client struct {
	hub *Hub
	conn *websocket.Conn
	//buffered channel of outbound messages
	send chan []byte
	Username string
	LobbyCode string
}

//pumps messages from websocket to hub
func (c *Client) readPump() {
	defer func() {
		c.hub.Unregister(c)
		_ = c.conn.Close()

		//update lobby via broadcast / attempt cleanup
		c.hub.BroadcastLobbyState(c.LobbyCode)
		_ = lobbyManager.RemoveIfEmpty(c.LobbyCode)
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ws read error: %v", err)
			}
			break
		}

		var incoming Message
		if err := json.Unmarshal(msgBytes, &incoming); err != nil {
			c.hub.Broadcast(msgBytes)
			continue
		}

		switch incoming.Type {
		case MessageTypeJoin:
			var jd JoinData
			if b, err := json.Marshal(incoming.Data); err == nil {
				_ = json.Unmarshal(b, &jd)
			}
			if jd.Username != "" {
				c.Username = jd.Username
			}
			c.hub.BroadcastLobbyState(c.LobbyCode)

			notice := Message{
				Type: MessageTypeChat,
				Data: map[string]string{"text": c.Username + " joined the lobby"},
			}
			c.hub.BroadcastMessage(notice)

		case MessageTypeLeave:
			c.hub.BroadcastMessage(incoming)

		default:
			c.hub.BroadcastMessage(incoming)
		}
	}
}

//writepump pumps messages from the hub to the websocket connection
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}

		case <-ticker.C:
			_= c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//makes a client, registers it with hub and starts the pumps
func NewClient(h *Hub, conn *websocket.Conn, lobbyCode string) *Client {
	c := &Client{
		hub: h,
		conn: conn,
		send: make(chan []byte, 256),
		Username: "",
		LobbyCode: lobbyCode,
	}

	//register with hub
	h.Register(c)
	go c.writePump()
	go c.readPump()

	return c
}