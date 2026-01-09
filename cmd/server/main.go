package main

import (
	"fmt"
	"time"

	"github.com/sitanshunandan/neurosync/internal/core/domain"
	"github.com/sitanshunandan/neurosync/internal/logic"
)

func main() {
	// 1. Setup the User (The Biology)
	now := time.Now()
	// Let's pretend the user wakes up at 7:00 AM today
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 7, 0, 0, 0, now.Location())

	userBio := domain.BioRhythm{
		UserID:       "user_001",
		WakeTime:     todayStart,
		SleepQuality: 0.90, // 90% sleep quality
	}

	// 2. Setup the Work (The To-Do List)
	// Notice we don't give them times, just "Cost"
	tasks := []domain.Task{
		{
			ID: "1", Title: "Deep Work: System Architecture",
			Cost:     domain.CognitiveLoad{Level: 9, Type: domain.LoadAnalytical},
			Duration: 2 * time.Hour,
		},
		{
			ID: "2", Title: "Reply to Emails",
			Cost:     domain.CognitiveLoad{Level: 3, Type: domain.LoadRote},
			Duration: 30 * time.Minute,
		},
		{
			ID: "3", Title: "Team Sync Meeting",
			Cost:     domain.CognitiveLoad{Level: 5, Type: domain.LoadAnalytical},
			Duration: 1 * time.Hour,
		},
		{
			ID: "4", Title: "Write Documentation",
			Cost:     domain.CognitiveLoad{Level: 7, Type: domain.LoadCreative},
			Duration: 1 * time.Hour,
		},
		{
			ID: "5", Title: "Fix Minor CSS Bug",
			Cost:     domain.CognitiveLoad{Level: 4, Type: domain.LoadRote},
			Duration: 30 * time.Minute,
		},
	}

	// 3. The Magic: Ask the Engine to Schedule
	fmt.Println("🧠 NeuroSync Engine | Optimizing Schedule...")
	schedule := logic.ScheduleTasks(userBio, tasks, todayStart)

	// 4. Print the Result
	printSchedule(schedule)
}

func printSchedule(s logic.Schedule) {
	fmt.Println("\n---------------------------------------------------------")
	fmt.Printf("OPTIMIZED SCHEDULE FOR %s\n", s.Date.Format("2006-01-02"))
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("%-10s | %-35s | %s\n", "TIME", "TASK", "ENERGY REQ")
	fmt.Println("---------------------------------------------------------")

	for _, item := range s.Items {
		timeStr := "Unscheduled"
		if item.FixedTime != nil {
			timeStr = item.FixedTime.Format("15:04")
		}

		// Visual indicator for difficulty
		difficulty := ""
		for i := 0; i < item.Cost.Level; i++ {
			difficulty += "⚡"
		}

		fmt.Printf("%-10s | %-35s | %s\n", timeStr, item.Title, difficulty)
	}
	fmt.Println("---------------------------------------------------------")
}
