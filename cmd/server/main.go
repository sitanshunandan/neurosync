package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sitanshunandan/neurosync/internal/adapters/handler"
	"github.com/sitanshunandan/neurosync/internal/adapters/repository"
)

func main() {
	// 1. Init Database (Creates 'neurosync.db' file)
	repo, err := repository.NewSQLiteRepository("./neurosync.db")
	if err != nil {
		fmt.Printf("🔥 Fatal: Could not connect to DB: %v\n", err)
		os.Exit(1)
	}

	// 2. Inject DB into Handler
	h := handler.NewHandler(repo)

	// 3. Setup Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/schedule", h.HandleSchedule)
	r.Get("/schedule/{userID}", h.HandleGetSchedule) // New GET route

	// 4. Start Server
	fmt.Println("🚀 NeuroSync Server (with SQLite) running on port 8080...")
	http.ListenAndServe(":8080", r)
}
