package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sitanshunandan/neurosync/internal/core/domain"
	"github.com/sitanshunandan/neurosync/internal/logic"
)

// CreateScheduleRequest is the JSON body the user sends us
type CreateScheduleRequest struct {
	UserID       string        `json:"user_id"`
	WakeTime     time.Time     `json:"wake_time"`     // Format: "2026-01-10T07:00:00Z"
	SleepQuality float64       `json:"sleep_quality"` // 0.0 to 1.0
	Tasks        []TaskRequest `json:"tasks"`
}

type TaskRequest struct {
	Title    string `json:"title"`
	Level    int    `json:"level"` // 1-10
	Type     string `json:"type"`  // "analytical", "creative", "rote"
	Duration int    `json:"duration_minutes"`
}

// Handler holds dependencies (like database connections, if we had any)
type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// HandleSchedule is the POST /schedule endpoint
func (h *Handler) HandleSchedule(w http.ResponseWriter, r *http.Request) {
	// 1. Decode JSON
	var req CreateScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 2. Map JSON Request to Domain Model (The "Adapter" pattern)
	// We convert the "web" structs to "internal" structs
	bio := domain.BioRhythm{
		UserID:       req.UserID,
		WakeTime:     req.WakeTime,
		SleepQuality: req.SleepQuality,
	}

	var tasks []domain.Task
	for _, t := range req.Tasks {
		// Default to Analytical if type is unknown
		loadType := domain.LoadAnalytical
		if t.Type == "creative" {
			loadType = domain.LoadCreative
		}
		if t.Type == "rote" {
			loadType = domain.LoadRote
		}

		tasks = append(tasks, domain.Task{
			Title:    t.Title,
			Cost:     domain.CognitiveLoad{Level: t.Level, Type: loadType},
			Duration: time.Duration(t.Duration) * time.Minute,
		})
	}

	// 3. Call the Logic (The Brain)
	// We use the request's WakeTime as the start of the day for simplicity
	schedule := logic.ScheduleTasks(bio, tasks, req.WakeTime)

	// 4. Encode Response (JSON)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}
