package notices

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrNoticeNotFound is returned when callers request a notice that does not exist.
	ErrNoticeNotFound = errors.New("notice not found")
	// ErrInvalidRequest is returned when a player request payload is incomplete.
	ErrInvalidRequest = errors.New("username and message are required")
)

// Notice represents a single notice/changelog entry to show in the UI.
type Notice struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Body        string    `json:"body"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"publishedAt"`
}

// PlayerRequest captures a suggestion or bug report submitted by a player.
type PlayerRequest struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

// PlayerRequestInput is the payload accepted from clients when submitting a request.
type PlayerRequestInput struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Sanitise trims whitespace and validates the payload.
func (in PlayerRequestInput) Sanitise() (PlayerRequestInput, error) {
	username := strings.TrimSpace(in.Username)
	message := strings.TrimSpace(in.Message)
	if username == "" || message == "" {
		return PlayerRequestInput{}, ErrInvalidRequest
	}
	return PlayerRequestInput{
		Username: username,
		Message:  message,
	}, nil
}
