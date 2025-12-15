package auth

// Credentials represents a minimal username/password login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
