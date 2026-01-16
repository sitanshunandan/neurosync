# NeuroSync: Circadian-Aware Task Scheduler

NeuroSync is a Go microservice that optimizes task scheduling by aligning workload difficulty with predicted biological alertness levels, based on the **Two-Process Model** of sleep regulation.

## Methodology

Unlike standard calendars that allocate time slots based only on availability, NeuroSync uses a bio-mathematical model to estimate cognitive capacity:

- **Process S (Homeostatic Sleep Drive):** Models fatigue accumulation relative to wake time.
- **Process C (Circadian Rhythm):** Models oscillations in alertness, including common variations like the afternoon energy dip.

The service accepts a task list with associated difficulty ratings and applies a greedy algorithm to:

- Assign **high-load tasks** to peak alertness windows
- Assign **low-load tasks** to lower-energy or recovery periods

## Scheduling Comparison

The table below compares a standard time-based schedule against the NeuroSync capacity-based schedule for a high-focus task (**System Design**) and a low-focus task (**Administrative Updates**).

| Time Window  | Standard Allocation        | NeuroSync Allocation                          |
|-------------|----------------------------|-----------------------------------------------|
| **10:00 AM** | Administrative Updates     | **System Design** (Peak Focus)                |
| **03:00 PM** | System Design              | **Administrative Updates** (Recovery)         |

## Technical Architecture

The application follows **Hexagonal Architecture** to separate business logic from external interfaces.

- **Language:** Go 1.25
- **Database:** SQLite
- **Containerization:** Docker (Alpine Linux)
- **Routing:** Chi

### Directory Structure

- `internal/core/domain`: Business logic entities and cognitive load models
- `internal/core/ports`: Input/output interfaces
- `internal/adapters`: Port implementations (HTTP handlers, SQL repositories)

## Usage

### Docker Build & Run

Build the image and run the service locally:

```bash
# 1. Build the image
docker build -t neurosync-backend .

# 2. Run the container (access at http://localhost:8080)
docker run -p 8080:8080 neurosync-backend
```

### API Endpoints
## 1. Generate Schedule (POST)

Submit a user's wake time and tasks to generate an optimized schedule.

Endpoint: POST /schedule

Body:
```bash
JSON

{
  "user_id": "test_user",
  "wake_time": "2026-01-14T07:00:00Z",
  "tasks": [
    { "title": "System Design", "duration": 120, "difficulty": 3 },
    { "title": "Email Catchup", "duration": 30, "difficulty": 1 }
  ]
}
```

## 2. Retrieve Schedule (GET)

Fetch the stored schedule for a specific user.

Endpoint: `GET /schedule/{user_id}`

Example: `GET /schedule/test_user`

