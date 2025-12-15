package notices

import (
	"context"
	"errors"
	"strings"
)

// Service exposes the notices/request operations needed by HTTP handlers.
type Service interface {
	ListNotices(ctx context.Context) ([]Notice, error)
	GetNotice(ctx context.Context, id string) (Notice, error)
	SubmitRequest(ctx context.Context, input PlayerRequestInput) (PlayerRequest, error)
	ListRequests(ctx context.Context) ([]PlayerRequest, error)
}

// Manager is the default service implementation backed by a Store.
type Manager struct {
	store Store
}

// NewService wires a Manager with the provided store.
func NewService(store Store) *Manager {
	return &Manager{store: store}
}

func (m *Manager) ListNotices(ctx context.Context) ([]Notice, error) {
	return m.store.ListNotices(ctx)
}

func (m *Manager) GetNotice(ctx context.Context, id string) (Notice, error) {
	if strings.TrimSpace(id) == "" {
		return Notice{}, ErrNoticeNotFound
	}
	return m.store.GetNotice(ctx, id)
}

func (m *Manager) SubmitRequest(ctx context.Context, input PlayerRequestInput) (PlayerRequest, error) {
	clean, err := input.Sanitise()
	if err != nil {
		return PlayerRequest{}, err
	}
	return m.store.SaveRequest(ctx, clean)
}

func (m *Manager) ListRequests(ctx context.Context) ([]PlayerRequest, error) {
	reqs, err := m.store.ListRequests(ctx)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

// Ensure Manager implements Service at compile time.
var _ Service = (*Manager)(nil)

// ErrInternal is returned when the service cannot fulfil a request.
var ErrInternal = errors.New("notices service unavailable")
