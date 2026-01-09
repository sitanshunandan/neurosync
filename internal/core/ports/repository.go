package ports

import "github.com/sitanshunandan/neurosync/internal/logic"

// SchedulerRepository defines how we store/retrieve schedules
type SchedulerRepository interface {
	Save(schedule logic.Schedule) error
	Get(userID string) (*logic.Schedule, error)
}
