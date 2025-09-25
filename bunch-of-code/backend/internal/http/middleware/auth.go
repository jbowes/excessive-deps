package middleware

import (
	"net/http"
)

// Auth is a simple auth middleware that checks for a static header.
// In real applications, replace with JWT/session validation.
func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-API-Key") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("missing X-API-Key"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
