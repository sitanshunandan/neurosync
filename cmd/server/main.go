package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sitanshunandan/neurosync/internal/adapters/handler"
)

func main() {
	// 1. Initialize Handlers
	h := handler.NewHandler()

	// 2. Setup Router
	r := chi.NewRouter()

	// Useful Middleware (Logging, Recovery)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 3. Define Routes
	r.Post("/schedule", h.HandleSchedule)

	// Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NeuroSync Engine is Active 🧠"))
	})

	// 4. Start Server
	fmt.Println("🚀 NeuroSync Server running on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
