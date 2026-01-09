package domain

import (
	"math"
	"time"
)

// CalculateCognitiveCapacity returns a score (0-100) for a specific time

func CalculateCognitiveCapacity(bio BioRhythm, targetTime time.Time) int {
	hoursAwake := targetTime.Sub(bio.WakeTime).Hours()

	// 1. Sleep Inertia: If they just woke up (< 1 hour), they are groggy.
	if hoursAwake < 1.0 {
		return int(30 * bio.SleepQuality)
	}

	// 2. Process S (Sleep Pressure):
	// Decays faster now (4.5 points/hr) to force a drop by evening.
	processS := 100.0 - (hoursAwake * 4.5)

	// 3. Process C (Circadian Drive):
	// High amplitude (25.0) to create a distinct Peak and Dip.
	// Shifted so peak is at ~3 hours awake, Dip at ~14 hours awake.
	circadianFluctuation := 25.0 * math.Sin((math.Pi/10.0)*(hoursAwake-2.0))

	totalCapacity := processS + circadianFluctuation

	// Apply Sleep Quality penalty
	totalCapacity *= bio.SleepQuality

	// Bounds checking
	if totalCapacity < 5 {
		return 5
	}
	if totalCapacity > 100 {
		return 100
	}

	return int(totalCapacity)
}
