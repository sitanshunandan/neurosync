package logic

import (
	"fmt"
	"sort"
	"time"

	"github.com/sitanshunandan/neurosync/internal/core/domain"
)

type Schedule struct {
	UserID string
	Date   time.Time
	Items  []domain.Task
}

func ScheduleTasks(bio domain.BioRhythm, tasks []domain.Task, dayStart time.Time) Schedule {
	// 1. Sort Tasks: Hardest first (Optimization strategy)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Cost.Level > tasks[j].Cost.Level
	})

	var scheduled []domain.Task

	// NEW: Keep track of the earliest free slot
	// We assume the user starts working at WakeTime + 1 hour (e.g., 8:00 AM)
	// (Or you can pass a specific 'WorkStartTime')
	currentTimePointer := bio.WakeTime.Add(1 * time.Hour)

	for _, task := range tasks {
		bestStartTime := time.Time{}
		foundSlot := false

		// Scan from the Current Pointer forward (up to 16h after wake)
		// We limit the search to avoid infinite loops
		searchLimit := bio.WakeTime.Add(16 * time.Hour)

		// Check every 30 mins from where we left off
		for t := currentTimePointer; t.Before(searchLimit); t = t.Add(30 * time.Minute) {

			// Ask the Brain: "Do we have energy here?"
			capacity := domain.CalculateCognitiveCapacity(bio, t)

			// Requirement: Task Level * 10 <= Capacity
			requiredCap := task.Cost.Level * 10

			if capacity >= requiredCap {
				bestStartTime = t
				foundSlot = true
				break // Found a spot!
			}
		}

		if foundSlot {
			t := bestStartTime
			task.FixedTime = &t
			scheduled = append(scheduled, task)

			// CRITICAL FIX: Advance the pointer!
			// The next task cannot start until this one finishes.
			currentTimePointer = bestStartTime.Add(task.Duration)
		} else {
			fmt.Printf("⚠️ Could not schedule task: %s (Req: %d, Best Cap: Low)\n", task.Title, task.Cost.Level*10)
		}
	}

	// Resort by Time for the display
	sort.Slice(scheduled, func(i, j int) bool {
		return scheduled[i].FixedTime.Before(*scheduled[j].FixedTime)
	})

	return Schedule{
		UserID: bio.UserID,
		Date:   dayStart,
		Items:  scheduled,
	}
}
