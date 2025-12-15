package storage

import (
	"context"
	"encoding/json"
	"testing"
)

func TestCreateLobby(t *testing.T) {
	store := NewInMemory()
	ctx := context.Background()

	t.Run("missing host", func(t *testing.T) {
		_, err := store.CreateLobby(ctx, "")
		if err != ErrMissingField {
			t.Fatalf("expected ErrMissingField, got %v", err)
		}
	})

	t.Run("creates lobby with host member", func(t *testing.T) {
		lobby, err := store.CreateLobby(ctx, "host")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if lobby.Code == "" || len(lobby.Code) != 6 {
			t.Fatalf("expected 6-char lobby code, got %q", lobby.Code)
		}
		if len(lobby.Members) != 1 || lobby.Members[0].Username != "host" || lobby.Members[0].Ready {
			t.Fatalf("unexpected members: %+v", lobby.Members)
		}
	})
}

func TestJoinAndGetLobby(t *testing.T) {
	store := NewInMemory()
	ctx := context.Background()

	lobby, _ := store.CreateLobby(ctx, "host")

	t.Run("missing fields", func(t *testing.T) {
		if _, err := store.JoinLobby(ctx, "", "user"); err != ErrMissingField {
			t.Fatalf("expected ErrMissingField, got %v", err)
		}
	})

	t.Run("lobby not found", func(t *testing.T) {
		if _, err := store.JoinLobby(ctx, "ZZZZZZ", "user"); err != ErrLobbyNotFound {
			t.Fatalf("expected ErrLobbyNotFound, got %v", err)
		}
	})

	t.Run("join success", func(t *testing.T) {
		got, err := store.JoinLobby(ctx, lobby.Code, "alice")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(got.Members) != 2 {
			t.Fatalf("expected 2 members, got %d", len(got.Members))
		}
	})

	t.Run("duplicate user rejected", func(t *testing.T) {
		if _, err := store.JoinLobby(ctx, lobby.Code, "alice"); err != ErrUserExists {
			t.Fatalf("expected ErrUserExists, got %v", err)
		}
	})

	t.Run("get lobby", func(t *testing.T) {
		got, err := store.GetLobby(ctx, lobby.Code)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got.Code != lobby.Code || len(got.Members) != 2 {
			t.Fatalf("unexpected lobby: %+v", got)
		}
	})
}

func TestSetReady(t *testing.T) {
	store := NewInMemory()
	ctx := context.Background()
	lobby, _ := store.CreateLobby(ctx, "host")
	_, _ = store.JoinLobby(ctx, lobby.Code, "alice")

	t.Run("updates flag", func(t *testing.T) {
		got, err := store.SetReady(ctx, lobby.Code, "alice", true)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !got.Members[1].Ready {
			t.Fatalf("expected alice to be ready, got %+v", got.Members[1])
		}
	})

	t.Run("missing user", func(t *testing.T) {
		if _, err := store.SetReady(ctx, lobby.Code, "ghost", true); err != ErrUserNotFound {
			t.Fatalf("expected ErrUserNotFound, got %v", err)
		}
	})
}

func TestLobbyCodeUniqueness(t *testing.T) {
	store := NewInMemory()
	ctx := context.Background()
	seen := make(map[string]struct{})

	for i := 0; i < 200; i++ {
		lobby, err := store.CreateLobby(ctx, "host")
		if err != nil {
			t.Fatalf("create lobby failed: %v", err)
		}
		if _, ok := seen[lobby.Code]; ok {
			t.Fatalf("duplicate lobby code generated: %s", lobby.Code)
		}
		seen[lobby.Code] = struct{}{}
	}
}

func TestSaveAndLoadResume(t *testing.T) {
	store := NewInMemory()
	ctx := context.Background()
	payload := json.RawMessage(`{"level":3,"hp":42}`)

	id, err := store.SaveResume(ctx, payload)
	if err != nil {
		t.Fatalf("save failed: %v", err)
	}
	got, err := store.LoadResume(ctx, id)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if string(got) != string(payload) {
		t.Fatalf("expected %s, got %s", string(payload), string(got))
	}
}
