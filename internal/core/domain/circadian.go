package domain

import (
	"math"
	"time"
)

// CalculateCognitiveCapacity returns a score (0-100) for a specific time
func CalculateCognitiveCapacity(bio BioRhythm, targetTime time.Time) int {
	// 1. Process S (Sleep Pressure): Linear decay since waking
	hoursAwake := targetTime.Sub(bio.WakeTime).Hours()
	if hoursAwake < 0 {
		return 0 // Hasn't woken up yet
	}

	// Homeostatic pressure increases with time awake.
	// Capacity drops linearly.
	processS := 100.0 - (hoursAwake * 5.5)

	// 2. Process C (Circadian Drive): Sine wave
	// We want peak at WakeTime + 4h.
	// This is a heuristic simulation of the Cortisol/Melatonin curves.
	circadianFluctuation := 15.0 * math.Sin((math.Pi/8.0)*(hoursAwake-1.0))

	totalCapacity := processS + circadianFluctuation

	// Penalize for poor sleep quality
	totalCapacity *= bio.SleepQuality

	if totalCapacity < 0 {
		return 0
	}
	if totalCapacity > 100 {
		return 100
	}

	return int(totalCapacity)
}
