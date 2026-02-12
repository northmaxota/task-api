package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mw "github.com/northmaxota/task-api/internal/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(mw.Logging)
	r.Use(middleware.Recoverer)

	r.Get("/health", HealthCheckHandler)

	return r
}

func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      NewRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
