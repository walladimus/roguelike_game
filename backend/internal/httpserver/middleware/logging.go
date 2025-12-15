package middleware

import (
	"log"
	"net/http"
	"time"
)

// loggingResponseWriter captures status codes for logging.
type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.status = code
	lrw.ResponseWriter.WriteHeader(code)
}

// RequestLogger logs method, path, status, and duration for each request.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(lrw, r)

		log.Printf("%s %s %d %s", r.Method, r.URL.Path, lrw.status, time.Since(start).Round(time.Millisecond))
	})
}
