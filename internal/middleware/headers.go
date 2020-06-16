package middleware

import (
	"net/http"
)

// Add additional security related headers
func Headers() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "SAMEORIGIN")
			w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Download-Options", "noopen")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

