package main

import (
	"log/slog" // Structured logging (Standard in Go 1.21+)
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "modernc.org/sqlite" // Pure Go SQLite driver

	"neurosync/internal/adapters/handler"
	"neurosync/internal/adapters/repository"
	"neurosync/internal/core/service"
)

func main() {
	// 1. Setup Structured Logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 2. Initialize Database (SQLite)
	// Note: In a cloud container, this file is ephemeral (wiped on restart).
	// For a portfolio demo, this is acceptable. For production, use Postgres.
	repo, err := repository.NewSQLiteRepository("neurosync.db")
	if err != nil {
		logger.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}

	// 3. Initialize Service & Handler
	svc := service.NewSchedulerService(repo)
	h := handler.NewHTTPHandler(svc)

	// 4. Setup Router & Middleware
	r := chi.NewRouter()

	// CLOUD REQUIREMENT: Request Logging
	// This lets you see "200 OK" or "500 ERROR" in your cloud dashboard logs.
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // Prevents crashes from taking down the server
	r.Use(middleware.Timeout(60 * time.Second))

	// 5. Define Routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NeuroSync Engine is Active 🧠"))
	})
	r.Post("/schedule", h.GenerateSchedule)
	r.Get("/schedule/{user_id}", h.GetSchedule)

	// 6. Start Server (Cloud Compatible)
	// Cloud platforms inject the PORT environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for localhost
	}

	logger.Info("Starting NeuroSync Server", "port", port, "env", "production")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
