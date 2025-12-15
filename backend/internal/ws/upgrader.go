package ws

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var allowedOrigins = loadAllowedOrigins()

// shared upgrader used by all ws handlers in this package.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// In dev, ALLOWED_ORIGINS may be empty: allow same-origin + localhost.
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true
		}
		if len(allowedOrigins) == 0 {
			return isLocalOrigin(origin)
		}
		for _, o := range allowedOrigins {
			if o == origin {
				return true
			}
		}
		return false
	},
}

func loadAllowedOrigins() []string {
	raw := os.Getenv("ALLOWED_ORIGINS")
	if raw == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimSpace(p); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func isLocalOrigin(origin string) bool {
	return strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "http://127.0.0.1:")
}
