package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/northmaxota/task-api/internal/config"
	mw "github.com/northmaxota/task-api/internal/middleware"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(mw.Logging)
	r.Use(middleware.Recoverer)

	r.Get("/health", HealthCheckHandler)

	return r
}

func NewServer(cfg config.Config) *http.Server {
	return &http.Server{
		Addr:         cfg.Port,
		Handler:      NewRouter(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
