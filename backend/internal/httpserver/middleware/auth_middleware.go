package middleware

import "net/http"

// Auth is a placeholder authentication middleware.
// Right now it just passes requests through unchanged.
//
// Later:
// - We'll check login/session tokens here.
// - We'll block requests if you're not allowed (e.g. starting lobby without being host).
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//no current logic
		next.ServeHTTP(w, r)
	})
}
