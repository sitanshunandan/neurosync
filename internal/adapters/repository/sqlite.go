package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sitanshunandan/neurosync/internal/logic"
	_ "modernc.org/sqlite" // Import the driver
)

type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository creates the DB file and the table if it doesn't exist
func NewSQLiteRepository(filePath string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite", filePath)
	if err != nil {
		return nil, err
	}

	// Create Table
	query := `
	CREATE TABLE IF NOT EXISTS schedules (
		user_id TEXT PRIMARY KEY,
		date TEXT,
		data TEXT
	);`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to init db: %w", err)
	}

	return &SQLiteRepository{db: db}, nil
}

// Save inserts or updates the schedule for a user
func (r *SQLiteRepository) Save(s logic.Schedule) error {
	// Serialize the Go struct to JSON
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	// UPSERT (Update if exists, Insert if new) - specific to SQLite syntax
	query := `
	INSERT INTO schedules (user_id, date, data) 
	VALUES (?, ?, ?)
	ON CONFLICT(user_id) DO UPDATE SET
		date=excluded.date,
		data=excluded.data;
	`

	_, err = r.db.Exec(query, s.UserID, s.Date.Format(time.RFC3339), string(data))
	return err
}

// Get retrieves the raw JSON and converts it back to Go struct
func (r *SQLiteRepository) Get(userID string) (*logic.Schedule, error) {
	row := r.db.QueryRow("SELECT data FROM schedules WHERE user_id = ?", userID)

	var dataStr string
	if err := row.Scan(&dataStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, err
	}

	var s logic.Schedule
	if err := json.Unmarshal([]byte(dataStr), &s); err != nil {
		return nil, err
	}

	return &s, nil
}
