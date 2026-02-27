package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/northmaxota/task-api/internal/config"
	httpServer "github.com/northmaxota/task-api/internal/http"
	"github.com/northmaxota/task-api/internal/storage"
	"github.com/northmaxota/task-api/internal/task"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	server := httpServer.NewServer(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		log.Println("Shutting down server...")
		shutdownContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownContext); err != nil {
			log.Printf("Server Shutdown Failed:%+v", err)
		}
	}()

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN environment variable is required")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	storage := storage.NewPostgresStorage(db)
	baseService := task.NewService(storage)
	service := task.NewLoggingService(baseService)
	handler := httpServer.NewHandler(service)
	router := httpServer.NewRouter(handler)
	server.Handler = router

	log.Printf("Starting server on %s\n", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
			return
		}
		log.Printf("Could not listen on %s: %v\n", cfg.Port, err)
	}

}
