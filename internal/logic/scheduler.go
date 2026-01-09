package logic

import (
	"sort"
	"time"

	"github.com/sitanshunandan/neurosync/internal/core/domain"
)

type Schedule struct {
	UserID string
	Date   time.Time
	Items  []domain.Task
}

// ScheduleTasks assigns time slots to tasks based on biological capacity
func ScheduleTasks(bio domain.BioRhythm, tasks []domain.Task, dayStart time.Time) Schedule {
	// 1. Sort Tasks: Hardest (Creative/Analytical) first.
	// We want to fit the "big rocks" into the jar first.
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Cost.Level > tasks[j].Cost.Level
	})

	var scheduled []domain.Task

	// 2. Iterate through tasks and find the best slot
	for _, task := range tasks {
		bestStartTime := time.Time{}
		bestScore := -1

		// Scan from WakeTime to +16 hours
		for h := 0; h < 16; h++ {
			slotTime := bio.WakeTime.Add(time.Duration(h) * time.Hour)

			// Ask the Brain: "How much energy do we have at this hour?"
			capacity := domain.CalculateCognitiveCapacity(bio, slotTime)

			// Requirement: Task Level * 10 <= Capacity
			// Level 8 task needs 80+ energy.
			requiredCap := task.Cost.Level * 10

			if capacity >= requiredCap {
				// Simple collision detection (MVP):
				// In a real app, we'd check if slot is already taken.
				// Here we just find the *earliest* matching slot.
				bestStartTime = slotTime
				bestScore = capacity
				break
			}
		}

		if bestScore != -1 {
			t := bestStartTime
			task.FixedTime = &t // Assign the time
			scheduled = append(scheduled, task)
		} else {
			// If no slot fits (e.g., too tired), we might log it or skip it
			// For now, let's skip appending it so we see what couldn't get done.
		}
	}

	// Resort by Time (so the printed schedule is chronological)
	sort.Slice(scheduled, func(i, j int) bool {
		return scheduled[i].FixedTime.Before(*scheduled[j].FixedTime)
	})

	return Schedule{
		UserID: bio.UserID,
		Date:   dayStart,
		Items:  scheduled,
	}
}
