package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware recovers from panics and returns a JSON error response
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("🔥 Panic: %v\n%s", rec, debug.Stack())

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(w).Encode(map[string]string{
					"error": "Internal Server Error",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
