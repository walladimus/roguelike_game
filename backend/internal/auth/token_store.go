package auth

import "sync"

// tokenStore holds ephemeral session tokens in memory.
type tokenStore struct {
	mu     sync.Mutex
	tokens map[string]string
}

func (t *tokenStore) generate(username string) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}
	t.mu.Lock()
	t.tokens[token] = username
	t.mu.Unlock()
	return token, nil
}
