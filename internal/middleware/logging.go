package middleware

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		reqID := GetRequestID(r)
		log.Printf("[reqID=%s] %s %s", reqID, r.Method, r.URL.RequestURI())

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		log.Printf("[reqID=%s] END %s %s %d %s", reqID, r.Method, r.URL.Path, rw.statusCode, duration)
	})
}
