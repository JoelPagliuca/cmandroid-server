package cmandroid

import (
	"log"
	"net/http"
	"time"
)

// Logger wraps a http.Handler with logging
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// BasicHeaders Adds the standard headers for a REST endpoint
func BasicHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inner.ServeHTTP(w, r)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	})
}
