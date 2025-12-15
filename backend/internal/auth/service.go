package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service exposes auth operations.
type Service interface {
	Register(ctx context.Context, creds Credentials) (string, error)
	Login(ctx context.Context, creds Credentials) (string, error)
	Validate(creds Credentials) error
}

// Manager provides register/login backed by a UserStore and in-memory tokens.
type Manager struct {
	store  UserStore
	tokens tokenStore
}

// NewService builds an auth manager with the provided user store.
func NewService(store UserStore) *Manager {
	return &Manager{
		store:  store,
		tokens: tokenStore{tokens: make(map[string]string)},
	}
}

func (s *Manager) Register(ctx context.Context, creds Credentials) (string, error) {
	if err := s.Validate(creds); err != nil {
		return "", err
	}

	hash, err := HashPassword(creds.Password)
	if err != nil {
		return "", err
	}
	if err := s.store.CreateUser(ctx, creds.Username, hash); err != nil {
		if errors.Is(err, ErrUserExists) {
			return "", err
		}
		return "", errors.New("failed to create user")
	}

	token, err := s.tokens.generate(creds.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Manager) Login(ctx context.Context, creds Credentials) (string, error) {
	if err := s.Validate(creds); err != nil {
		return "", err
	}

	hash, err := s.store.GetUserHash(ctx, creds.Username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", errors.New("failed to fetch user")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.tokens.generate(creds.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *Manager) Validate(creds Credentials) error {
	if creds.Username == "" || creds.Password == "" {
		return errors.New("username and password required")
	}
	if len(creds.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
