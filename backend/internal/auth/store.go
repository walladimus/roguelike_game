package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"sync"
)

var (
	ErrUserExists   = errors.New("username already exists")
	ErrUserNotFound = errors.New("user not found")
)

// UserStore persists auth users.
type UserStore interface {
	CreateUser(ctx context.Context, username, passwordHash string) error
	GetUserHash(ctx context.Context, username string) (string, error)
}

// InMemoryUserStore keeps users in memory (dev fallback).
type InMemoryUserStore struct {
	mu    sync.Mutex
	users map[string]string
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{users: make(map[string]string)}
}

func (s *InMemoryUserStore) CreateUser(_ context.Context, username, passwordHash string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[username]; exists {
		return ErrUserExists
	}
	s.users[username] = passwordHash
	return nil
}

func (s *InMemoryUserStore) GetUserHash(_ context.Context, username string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	hash, ok := s.users[username]
	if !ok {
		return "", ErrUserNotFound
	}
	return hash, nil
}

// PostgresUserStore persists users in Postgres.
type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{db: db}
}

func (s *PostgresUserStore) CreateUser(ctx context.Context, username, passwordHash string) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO users (username, password_hash)
		VALUES ($1, $2)
	`, username, passwordHash)
	if err != nil {
		// simple duplicate detection using constraint error text
		if isUniqueViolation(err) {
			return ErrUserExists
		}
		return err
	}
	return nil
}

func (s *PostgresUserStore) GetUserHash(ctx context.Context, username string) (string, error) {
	var hash string
	err := s.db.QueryRowContext(ctx, `SELECT password_hash FROM users WHERE username = $1`, username).Scan(&hash)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrUserNotFound
	}
	return hash, err
}

// isUniqueViolation tries to detect Postgres unique constraint errors without pgconn.
func isUniqueViolation(err error) bool {
	// pgx stdlib exposes errors with Code() when using pgconn's pgerror, but here we just fallback to string match.
	return err != nil && (errors.Is(err, ErrUserExists) || strings.Contains(strings.ToLower(err.Error()), "unique"))
}
