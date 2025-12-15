package ws

// read only view of a lobby for debugging
type LobbyInfo struct {
	Code    string   `json:"code"`
	Players []string `json:"players"`
	Count   int      `json:"count"`
}

func (m *LobbyManager) Snapshot() []LobbyInfo {
	out := make([]LobbyInfo, 0)

	m.ForEach(func(code string, h *Hub) {
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

		out = append(out, LobbyInfo{
			Code:    code,
			Players: players,
			Count:   len(players),
		})
	})

	return out
}

// exposes snapshot using package-level manager. prevents httpserver from depending on ws internals.
func SnapshotLobbies() []LobbyInfo {
	return lobbyManager.Snapshot()
}
