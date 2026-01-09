package domain

import (
	"time"
)

// CognitiveLoad represents the "Biological Cost" of a task
type CognitiveLoad struct {
	Level int      // 1 (Low) to 10 (High)
	Type  LoadType // Creative, Analytical, Rote
}

type LoadType string

const (
	LoadCreative   LoadType = "creative"   // High Dopamine demand
	LoadAnalytical LoadType = "analytical" // High Norepinephrine demand
	LoadRote       LoadType = "rote"       // Low metabolic demand
)

type Task struct {
	ID        string
	Title     string
	Cost      CognitiveLoad
	Duration  time.Duration
	Deadline  time.Time
	FixedTime *time.Time // nil if floating
}

// BioRhythm represents the user's current chemical state
type BioRhythm struct {
	UserID       string
	WakeTime     time.Time
	SleepTime    time.Time // Last night's sleep start
	SleepQuality float64   // 0.0 - 1.0
}
